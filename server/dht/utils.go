package dht

func (s *Server) isOnlyNodeInNetwork() bool {
	return s.predecessor == nil && s.id.Equals(s.Successor())
}
