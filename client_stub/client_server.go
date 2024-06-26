package client_stub

import (
	"bytes"
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io"
	"net"
	"net/http"
	"securechat-server/client_stub/grpc"
	"securechat-server/globals"
	"securechat-server/server/dht/records"
	requests "securechat-server/server/types"
	"strings"
	"sync"
)

type GRPCServer struct {
	mu sync.Mutex
	grpc.UnimplementedClientServerCommsServer
	Requests   chan<- requests.Request
	Response   <-chan records.Record
	Challenges *Challenges
}

// SignUp
/*
Sign up takes the information passed by the client.
It generates a cryptographically secure challenge and sends it back to the client, encrypted with the client's public key.
It then adds the challenge to the map of challenges, so that it can then be verified later once the client responds.
*/
func (s *GRPCServer) SignUp(ctx context.Context, request *grpc.SignUpRequest) (*grpc.SignUpResponse, error) {
	// TODO: Check if user already exists

	// Get IP address of the client
	ip := getClientIP(ctx)
	// Create record struct
	record := records.Record{
		Username:       request.Username,
		Address:        ipToUint32(ip),
		PublicKeyChat:  request.PublicKeyChat,
		PublicKeyLogin: request.PublicKeyLogin,
	}

	randChal, err := GenerateRandomChallenge()
	if err != nil {
		panic("Failed to generate challenge")
	}

	// Create challenge struct
	chal := Challenge{
		C: randChal,
		R: record,
	}

	publicKeyLogin, err := x509.ParsePKIXPublicKey(request.PublicKeyLogin)
	if err != nil {
		panic("Failed to parse RSA public key")
	}

	encryptedChal, err := encryptUint64(chal.C, publicKeyLogin.(*rsa.PublicKey))
	if err != nil {
		return nil, err
	}

	// Add challenge to map
	s.Challenges.Add(ip, chal)

	return &grpc.SignUpResponse{
		Challenge: encryptedChal,
	}, nil
}

// SignUpChallengeResponse
/*
This is called when the client is responding to a signup challenge sent by the server.
The server gets the clients IP address then retrieves the challenge from the map of challenges.
It then uses the record retrieved from the map to decrypt the challenge using the client's public key and verify it.

If the challenge is verified, the server adds the record to the DHT and sends a success response to the client.
If the challenge is not verified, the server sends a failure response to the client.
*/
func (s *GRPCServer) SignUpChallengeResponse(ctx context.Context, request *grpc.SignUpChallengeResponseRequest) (*grpc.SignUpChallengeResponseResponse, error) {
	// Get IP address of the client
	ip := getClientIP(ctx)

	// Get challenge from map
	chal := s.Challenges.Get(ip)
	s.Challenges.Remove(ip)

	// Decrypt challenge response with public key and verify
	publicKeyLogin, err := x509.ParsePKIXPublicKey(chal.R.PublicKeyLogin)
	if err != nil || !verifySignature(uint64ToBytes(chal.C-1), request.ChallengeResponse, publicKeyLogin.(*rsa.PublicKey)) {
		return nil, status.Error(codes.PermissionDenied, "Invalid")
	}

	// If challenge is verified, add record to DHT
	req := requests.Request{
		Type:   requests.PUT,
		Record: chal.R,
	}
	s.mu.Lock()
	s.Requests <- req
	s.mu.Unlock()

	return &grpc.SignUpChallengeResponseResponse{}, nil
}

// Login
/*
Login verifies the digital signature of the client using the public key stored in the record.
If the digital signature is verified, the server updates the IP address in the record and sends a success response to the client.
If the digital signature is not verified, the server sends an error response to the client.
*/
func (s *GRPCServer) Login(ctx context.Context, request *grpc.LoginRequest) (*grpc.LoginResponse, error) {
	// Get record from server
	user := s.getUser(request.DigitalSignature.Username)

	// Verify digital signature using public key stored in record
	publicKeyLogin, err := x509.ParsePKIXPublicKey(user.PublicKeyLogin)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "Invalid signature")
	}
	message := []byte(strings.Join([]string{request.Username, uint32ToIp(request.Address).String()}, ";"))
	verified := verifySignature(message, request.DigitalSignature.Signature, publicKeyLogin.(*rsa.PublicKey))

	// if not verified, send error
	if !verified {
		return nil, status.Error(codes.PermissionDenied, "Invalid signature")
	}

	// If user doesn't exist, send error
	if user.Username == "" {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	// challenge IP so the client proves it exists
	c_int, err := GenerateRandomChallenge()
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, "Error generating challenge")
	}

	c, err := encryptUint64(c_int, publicKeyLogin.(*rsa.PublicKey))

	url := fmt.Sprintf("http://%s:6500/challenge/%d", uint32ToIp(request.Address), request.ChallengeNonce)
	hreq, err := http.NewRequest("POST", url, bytes.NewBuffer(c))
	if err != nil {
		return nil, status.Error(codes.NotFound, "Error sending challenge to client")
	}

	// Set the request content type
	hreq.Header.Set("Content-Type", "application/octet-stream")

	// Create an HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(hreq)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Error sending challenge to client")
	}
	defer resp.Body.Close()

	// Read the response body
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Error reading challenge response")
	}

	verified = verifySignature(uint64ToBytes(c_int-1), responseBytes, publicKeyLogin.(*rsa.PublicKey))
	if !verified {
		return nil, status.Error(codes.PermissionDenied, "Invalid signature")
	}

	// Update IP address in record
	user.Address = request.Address

	// Create request struct
	req := requests.Request{
		Type:   requests.PUT,
		Record: user,
	}

	// Send request to server
	s.mu.Lock()
	s.Requests <- req
	s.mu.Unlock()

	// Send success response
	return &grpc.LoginResponse{}, nil
}

// FindUser
/*
FindUser takes a username and returns the record associated with that username.
*/
func (s *GRPCServer) FindUser(ctx context.Context, request *grpc.FindUserRequest) (*grpc.FindUserResponse, error) {
	// Create record struct
	record := s.getUser(request.Username)
	sender := s.getUser(request.DigitalSignature.Username)

	// Verify digital signature using public key stored in record
	publicKeyLogin, err := x509.ParsePKIXPublicKey(sender.PublicKeyLogin)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "Bad public key for signature")
	}
	message := []byte(request.Username)
	verified := verifySignature(message, request.DigitalSignature.Signature, publicKeyLogin.(*rsa.PublicKey))

	// if not verified, send error
	if !verified {
		return nil, status.Error(codes.PermissionDenied, "Invalid signature")
	}

	if record.Username == "" {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	return &grpc.FindUserResponse{
		Username:      record.Username,
		Address:       record.Address,
		PublicKeyChat: record.PublicKeyChat,
	}, nil
}

func NewGRPCClientServer(requests chan<- requests.Request, response <-chan records.Record) {
	server := GRPCServer{Requests: requests, Response: response, Challenges: NewChallenges()}

	tlsCert, err := tls.LoadX509KeyPair("certs/server-cert.pem", "certs/server-key.pem")
	if err != nil {
		panic(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		ClientAuth:   tls.NoClientCert,
	}

	s := ggrpc.NewServer(ggrpc.Creds(credentials.NewTLS(config)))

	grpc.RegisterClientServerCommsServer(s, &server)

	lis, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", globals.ServerAddress, globals.ClientPort))
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
