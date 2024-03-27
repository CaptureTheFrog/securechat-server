package dht

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	grpc2 "securechat-server/server/dht/grpc"
	pb "securechat-server/server/dht/monitor"
	"securechat-server/server/dht/types"
	"time"
)

func (s *Server) Maintain() {
	// Create client to monitor server
	conn, err := grpc.Dial("127.0.0.1:54000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	client := pb.NewMonitorCommsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for {
		time.Sleep(10 * time.Second)
		s.checkPredecessor()
		s.checkSuccessor()

		// print predecessor and successor
		s.mu.RLock()
		pred := s.predecessor
		succ := s.Successor()
		s.mu.RUnlock()

		var predName string
		var succName string
		address := s.id.Name

		if pred == nil {
			predName = "null"
		} else {
			predName = pred.Name
		}

		if succ == nil {
			succName = "null"
		} else {
			succName = succ.Name
		}

		// Inform monitor server
		Inform(client, ctx, address, predName, succName)

		fmt.Printf("\nPredecessor: %s\n", predName)
		fmt.Printf("Successor: %s\n", succName)
		s.printRecords()
	}
}

func (s *Server) printRecords() {
	fmt.Printf("Records:\n")
	for _, v := range s.records.GetAll() {
		fmt.Printf("\t")
		v.Print()
	}
}

// Inform
func Inform(client pb.MonitorCommsClient, ctx context.Context, addr string, pred string, succ string) {
	_, _ = client.Inform(ctx, &pb.Info{
		Address:     addr,
		Predecessor: pred,
		Successor:   succ,
	})
}

// checkPredecessor
/*
	Checks if the predecessor P is still the predecessor
	Asks P for its successor
	If the successor of P (S) is not this server, then:
		- If S is between P and this server
			P -> S -> This server

			- Update the successor of S to this server
			- Update the predecessor of this server to S
		- If this server is between P and S
			P -> This server -> S

			- Update the successor of P to this server
			- Update the predecessor of the S to this server
			- Update the successor of this server to S
*/
func (s *Server) checkPredecessor() {
	thisPredecessor := s.Predecessor()

	// If there isn't a predecessor
	if thisPredecessor == nil {
		return
	}

	// Create client
	client, err := CreateGRPCClient(thisPredecessor.Name)
	if err != nil {
		s.UpdatePredecessor(nil)
	}

	// Ask predecessor for its successor
	succ, err := client.GetSuccessor(client.Ctx, thisPredecessor.ToGRPC())
	client.Cancel()

	if err != nil {
		s.UpdatePredecessor(nil)
	}

	successor := types.IDFromGRPC(succ)
	// If the successor of the predecessor is not this server
	if !successor.Equals(s.id) {

		// If the successor comes between predecessor and this server
		if successor.IsBetween(thisPredecessor, s.id) {
			client.Cancel()

			// "Successor" is now this server's predecessor
			// Create client to update "successor"
			client, err = CreateGRPCClient(successor.Name)
			if err != nil {
				return
			}

			// Update "successor's" successor
			_, err = client.ChangeSuccessor(client.Ctx, &grpc2.ChangeSuccessor{
				Id:           successor.ToGRPC(),
				NewSuccessor: s.id.ToGRPC(),
			})
			client.Cancel()
			if err != nil {
				return
			}

			// Update this server's predecessor
			s.UpdatePredecessor(successor)
		} else {
			// This server comes between the predecessor and "successor"

			// Update predecessor
			_, err = client.ChangeSuccessor(client.Ctx, &grpc2.ChangeSuccessor{
				Id:           s.predecessor.ToGRPC(),
				NewSuccessor: s.id.ToGRPC(),
			})
			client.Cancel()
			if err != nil {
				return
			}

			// If "successor" comes between this and its successor
			if successor.IsBetween(s.id, s.Successor()) {
				// Update "successor"
				client, err = CreateGRPCClient(successor.Name)
				if err != nil {
					return
				}

				// Update "successor's" predecessor
				_, err = client.ChangePredecessor(client.Ctx, &grpc2.ChangePredecessor{
					Id:             successor.ToGRPC(),
					NewPredecessor: s.id.ToGRPC(),
				})
				client.Cancel()
				if err != nil {
					return
				}

				// Update this server's successor
				s.UpdateSuccessor(successor)
			}
		}
	}
}

// checkSuccessor
/*
	Checks if the successor S is still the successor of this server
	Asks S for its predecessor
	If the predecessor of S (P) is not this server, then:
		- If P is between this server and S
			This server -> P -> S

			- Update the predecessor of P to this server
			- Update the successor of this server to P
		- If this server is between P and S
			P -> This server -> S

			- Update the successor of P to this server
			- Update the predecessor of this server to P
*/
func (s *Server) checkSuccessor() {
	thisSuccessor := s.Successor()

	// If this server is the only one in the network
	if s.isOnlyNodeInNetwork() {
		return
	}

	// Create client
	client, err := CreateGRPCClient(thisSuccessor.Name)
	if err != nil {
		// TODO: Move to next successor
		return
	}

	// Ask successor for its predecessor
	pred, err := client.GetPredecessor(client.Ctx, thisSuccessor.ToGRPC())
	client.Cancel()
	if err != nil {
		// TODO: Move to next successor
		return
	}

	predecessor := types.IDFromGRPC(pred)

	// If the predecessor of the successor is not this server
	if !predecessor.Equals(s.id) {
		// If the predecessor comes between this server and the successor
		if predecessor.IsBetween(s.id, thisSuccessor) {
			// "Predecessor" is now this server's successor
			// Create client to update "predecessor"
			client, err = CreateGRPCClient(predecessor.Name)
			if err != nil {
				return
			}
			_, err = client.ChangePredecessor(client.Ctx, &grpc2.ChangePredecessor{
				Id:             predecessor.ToGRPC(),
				NewPredecessor: s.id.ToGRPC(),
			})
			client.Cancel()
			if err != nil {
				return
			}

			// Update successor on this server
			s.UpdateSuccessor(predecessor)

		} else {
			// This server comes between the predecessor and the successor
			// Update "predecessor" to have this server as successor
			// Create client
			client, err = CreateGRPCClient(predecessor.Name)
			if err != nil {
				return
			}
			_, err = client.ChangeSuccessor(client.Ctx, &grpc2.ChangeSuccessor{
				Id:           predecessor.ToGRPC(),
				NewSuccessor: s.id.ToGRPC(),
			})
			client.Cancel()
			if err != nil {
				return
			}

			// Update successor's predecessor
			// Create client
			client, err = CreateGRPCClient(thisSuccessor.Name)
			if err != nil {
				return
			}
			_, err = client.ChangePredecessor(client.Ctx, &grpc2.ChangePredecessor{
				Id:             thisSuccessor.ToGRPC(),
				NewPredecessor: s.id.ToGRPC(),
			})
			client.Cancel()
			if err != nil {
				return
			}

			// This server's predecessor
			s.UpdatePredecessor(predecessor)
		}
	}
}
