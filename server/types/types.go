package types

type RequestType int

const (
	GET RequestType = iota
	PUT
)

// Request is the type of request that can be made to the internal DHT server.
type Request struct {
	Type   RequestType
	Record Record
}

// Record is the type of record that is stored in the internal DHT server.
type Record struct {
	Username  string
	Address   string
	PublicKey string
}
