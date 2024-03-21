package client_stub

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"fmt"
	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"math/big"
	"net"
	"securechat-server/client_stub/grpc"
	"securechat-server/globals"
	"securechat-server/server/dht/records"
	requests "securechat-server/server/types"
	"strconv"
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
	// Get IP address of the client
	ip := getClientIP(ctx)
	// Create record struct
	record := records.Record{
		Username:       request.Username,
		Address:        ipToUint32(ip),
		PublicKeyChat:  request.PublicKeyChat,
		PublicKeyLogin: request.PublicKeyLogin,
	}

	// Create challenge struct
	chal := Challenge{
		C: GenerateRandomChallenge(),
		R: record,
	}

	publicKeyLogin, err := x509.ParsePKCS1PublicKey(request.PublicKeyLogin)
	if err != nil {
		panic("Failed to parse RSA public key")
	}

	encryptedChal, err := encryptUint64(chal.C, publicKeyLogin)
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

	// Decrypt challenge response with public key
	// TODO: Decrypt challenge with public key
	decryptedChal := request.ChallengeResponse

	// Verify challenge
	if chal.C != decryptedChal+1 {
		// Throw error if challenge is not correct
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
	username := request.Username

	// Get record from server
	user := s.getUser(username)

	// Verify digital signature using public key stored in record
	publicKeyLogin, err := x509.ParsePKCS1PublicKey(user.PublicKeyLogin)
	if err != nil {
		panic("Failed to parse RSA public key")
	}
	message := []byte(strings.Join([]string{request.Username, uint32ToIp(request.Address).String()}, ";"))
	verified := verifySignature(message, request.DigitalSignature.Signature, publicKeyLogin)

	// if not verified, send error
	if !verified {
		return nil, status.Error(codes.PermissionDenied, "Invalid signature")
	}

	// If user doesn't exist, send error
	if user.Username == "" {
		return nil, status.Error(codes.NotFound, "User not found")
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

	// TODO: Verify digital signature

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
	s := ggrpc.NewServer()

	grpc.RegisterClientServerCommsServer(s, &server)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", globals.ServerAddress, globals.ClientPort))
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

// ipToUint32
// Convert net.IP to uint32
func ipToUint32(addr net.Addr) uint32 {
	ip, ok := addr.(*net.IPAddr)
	if !ok {
		panic("Invalid net.Addr type, must be net.IPAddr")
	}

	ipString := ip.IP.String()

	octets := strings.Split(ipString, ".")
	var result uint32
	for i, octet := range octets {
		octetVal, err := strconv.Atoi(octet)
		if err != nil {
			panic(err)
		}
		result |= octetVal << ((3 - i) * 8)
	}

	return result
}

// uint32ToIp
// Convert uint32 to net.IP
func uint32ToIp(value uint32) net.IP {
	ip := make(net.IP, 4)
	for i := 0; i < 4; i++ {
		shift := uint(8 * (3 - i))
		octet := byte((value >> shift) & 0xFF)
		ip[i] = octet
	}
	return ip
}

func (s *GRPCServer) getUser(username string) records.Record {
	// Create record struct
	record := records.Record{
		Username:       username,
		Address:        0,
		PublicKeyChat:  make([]byte, 0),
		PublicKeyLogin: make([]byte, 0),
	}

	// Create request struct
	req := requests.Request{
		Type:   requests.GET,
		Record: record,
	}

	s.mu.Lock()
	s.Requests <- req
	record = <-s.Response
	s.mu.Unlock()

	return record
}

func getClientIP(ctx context.Context) net.Addr {
	peerAddr, ok := peer.FromContext(ctx)
	if ok {
		return peerAddr.Addr
	}
	return nil
}

func verifySignature(message []byte, signature []byte, publicKey *rsa.PublicKey) bool {
	hashed := sha256.Sum256(message)
	err := rsa.VerifyPSS(publicKey, crypto.SHA256, hashed[:], signature, nil)
	if err != nil {
		return false
	}
	return true
}

func encryptUint64(value uint64, publicKey *rsa.PublicKey) ([]byte, error) {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, new(big.Int).SetUint64(value).Bytes(), nil)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}
