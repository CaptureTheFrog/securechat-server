package client_stub

import (
	"crypto/rand"
	"math/big"
	"net"
	"securechat-server/server/dht/records"
	"sync"
)

type Challenges struct {
	mu         sync.Mutex
	challenges map[net.Addr]Challenge
}

type Challenge struct {
	C uint64
	R records.Record
}

func NewChallenges() *Challenges {
	return &Challenges{challenges: make(map[net.Addr]Challenge)}
}

func (c *Challenges) Add(addr net.Addr, chal Challenge) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.challenges[addr] = chal
}

func (c *Challenges) Get(addr net.Addr) Challenge {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.challenges[addr]
}

func (c *Challenges) Remove(addr net.Addr) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.challenges, addr)
}

// GenerateRandomChallenge
// Generates a random 64 bit cryptographically secure challenge
func GenerateRandomChallenge() (uint64, error) {
	randomNumber, err := rand.Int(rand.Reader, new(big.Int).SetUint64((^uint64(0))-1))
	if err != nil {
		return 0, err
	}
	return randomNumber.Uint64() + 1, nil
}
