package globals

import "net"

var (
	GRPCPort      int
	Method        string
	JoinAddress   string
	ServerAddress string
	ClientPort    int
)

func FindFreePort() (int, net.IP, error) {
	// Listen on port 0 to let the OS choose an available port
	listener, err := net.Listen("tcp4", "0.0.0.0:0")
	if err != nil {
		return 0, nil, err
	}
	defer listener.Close()

	// Get the address from the listener and extract the port
	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port, addr.IP, nil
}
