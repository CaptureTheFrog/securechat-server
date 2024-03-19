package dht

import (
	"fmt"
	"time"
)

func (s *Server) Maintain() {
	for {
		// print predecessor and successor
		s.mu.Lock()
		pred := s.predecessor
		succ := s.Successor()
		s.mu.Unlock()

		var predName string
		var succName string

		if pred == nil {
			predName = "null"
		} else {
			predName = pred.name
		}

		if succ == nil {
			succName = "null"
		} else {
			succName = succ.name
		}

		fmt.Printf("\nPredecessor: %s\n", predName)
		fmt.Printf("Successor: %s\n", succName)

		// sleep for 5 seconds
		time.Sleep(10 * time.Second)
	}
}
