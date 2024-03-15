package main

import (
	. "securechat-server/server"
)

func main() {
	s := NewServer("localhost:50051")

	record := s.Get("test")

	println(record.Address)
}
