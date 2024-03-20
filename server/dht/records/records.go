package records

import (
	"fmt"
	"securechat-server/server/dht/types"
)

type Records struct {
	records map[*types.ID]Record
}

func NewRecords() *Records {
	return &Records{records: make(map[*types.ID]Record)}
}

func (r *Records) Add(record Record) {
	id := types.NewID(record.Username)
	r.records[id] = record
}

func (r *Records) Get(id *types.ID) Record {
	// loop through records and find the one with the same id
	for k, v := range r.records {
		if k.Equals(id) {
			return v
		}
	}

	return *NewRecord()
}

func (r *Records) GetAll() map[*types.ID]Record {
	return r.records
}

// Record is the type of record that is stored in the internal DHT server.
type Record struct {
	Username       string
	Address        uint32
	PublicKeyLogin []byte
	PublicKeyChat  []byte
}

func (r *Record) Print() {
	fmt.Printf("Username: %s\tAddress: %s\tPublic Key Login: %v\tPublic Key Chat: %v\n",
		r.Username, r.Address, r.PublicKeyLogin, r.PublicKeyChat)
}

func NewRecord() *Record {
	return &Record{
		Username:       "",
		Address:        0,
		PublicKeyLogin: make([]byte, 0),
		PublicKeyChat:  make([]byte, 0),
	}
}
