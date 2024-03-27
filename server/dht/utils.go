package dht

import "strings"

func (s *Server) isOnlyNodeInNetwork() bool {
	return s.Predecessor() == nil && s.id.Equals(s.Successor())
}

func EqualAddresses(a string, b string) bool {
	// Split addresses at colon
	aSplit := strings.Split(a, ":")
	bSplit := strings.Split(b, ":")

	// Compare ports
	if aSplit[1] != bSplit[1] {
		return false
	}

	// Compare IP addresses
	// If the IP addresses aren't the same, return false
	if aSplit[0] != bSplit[0] {
		// 127.0.0.1 and 0.0.0.0 are the same
		if (aSplit[0] == "127.0.0.1" || aSplit[0] == "0.0.0.0") &&
			(bSplit[0] == "127.0.0.1" || bSplit[0] == "0.0.0.0") {
			return true
		}

		return false
	}

	return true
}
