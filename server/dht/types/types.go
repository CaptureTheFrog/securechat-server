package types

import (
	"bytes"
	"crypto/sha1"
	"securechat-server/server/dht/grpc"
)

type ID struct {
	// ID is the unique identifier for a node in the DHT.
	ID   [20]byte
	Name string
}

func NewID(name string) *ID {
	// Sha1 hash the name to get a unique identifier.
	hasher := sha1.New()
	hasher.Write([]byte(name))
	hashedBytes := [20]byte(hasher.Sum(nil))

	// Make a copy to separate it from the hasher
	copyBytes := make([]byte, 20)
	copy(copyBytes, hashedBytes[:])

	return &ID{ID: [20]byte(copyBytes), Name: name}
}

func IDFromGRPC(id *grpc.ID) *ID {
	if *id.Address == "" {
		return &ID{}
	}

	return &ID{ID: [20]byte(id.Id), Name: *id.Address}
}

func (id *ID) Equals(otherId *ID) bool {
	return bytes.Equal(id.ID[:], otherId.ID[:])
}

func (id *ID) IsBetween(start *ID, end *ID) bool {
	// If the start ID is less than the end ID,
	// then the ID is between them if it is greater than the start and less than the end.
	if bytes.Compare(start.ID[:], end.ID[:]) < 0 {
		return bytes.Compare(start.ID[:], id.ID[:]) < 0 && bytes.Compare(id.ID[:], end.ID[:]) < 0
	}

	// If the start ID is greater than the end ID,
	// then the ID is between them if it is greater than the start or less than the end.
	if bytes.Compare(start.ID[:], end.ID[:]) > 0 {
		return bytes.Compare(start.ID[:], id.ID[:]) < 0 || bytes.Compare(id.ID[:], end.ID[:]) < 0
	}

	// If they are equal
	return false
}

func (id *ID) ToGRPC() *grpc.ID {
	return &grpc.ID{Address: &id.Name, Id: id.ID[:]}
}

type FingerTable struct {
	entries map[int]*ID
}

func NewFingerTable() *FingerTable {
	return &FingerTable{entries: make(map[int]*ID)}
}

func (ft *FingerTable) AddEntry(index int, id *ID) {
	ft.entries[index] = id
}

func (ft *FingerTable) GetEntry(index int) *ID {
	return ft.entries[index]
}
