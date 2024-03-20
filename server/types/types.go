package types

import "securechat-server/server/dht/records"

type RequestType int

const (
	GET RequestType = iota
	PUT
)

// Request is the type of request that can be made to the internal DHT server.
type Request struct {
	Type   RequestType
	Record records.Record
}
