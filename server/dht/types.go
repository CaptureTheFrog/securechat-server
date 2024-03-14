package dht

import (
	"bytes"
	"crypto/sha1"
)

type ID struct {
	// ID is the unique identifier for a node in the DHT.
	ID   [20]byte
	name string
}

func NewID(name string) *ID {
	// Sha1 hash the name to get a unique identifier.
	hasher := sha1.New()
	hasher.Write([]byte(name))
	bytes := [20]byte(hasher.Sum(nil))

	// Make a copy to separate it from the hasher
	copyBytes := make([]byte, 20)
	copy(bytes[:], copyBytes)

	return &ID{ID: [20]byte(copyBytes), name: name}
}

func (id *ID) equals(otherId *ID) bool {
	return bytes.Equal(id.ID[:], otherId.ID[:])
}

type FingerTable struct {
	entries map[int]*ID
}

func NewFingerTable() *FingerTable {
	return &FingerTable{entries: make(map[int]*ID)}
}
