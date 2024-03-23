package client_stub

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/binary"
	"google.golang.org/grpc/peer"
	"math/big"
	"net"
	"securechat-server/server/dht/records"
	requests "securechat-server/server/types"
	"strconv"
	"strings"
)

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

// ipToUint32
// Convert net.IP to uint32
func ipToUint32(addr *net.TCPAddr) uint32 {
	ipString := addr.IP.String()

	octets := strings.Split(ipString, ".")
	var result uint32
	for i, octet := range octets {
		octetVal, err := strconv.Atoi(octet)
		if err != nil {
			panic(err)
		}
		result |= uint32(octetVal) << ((3 - i) * 8)
	}

	return result
}

// Function to convert uint64 to bytes
func uint64ToBytes(value uint64) []byte {
	bytes := make([]byte, 8) // 8 bytes for uint64
	binary.BigEndian.PutUint64(bytes, value)
	return bytes
}
