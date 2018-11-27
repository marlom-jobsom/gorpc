package network

import (
	"fmt"
	"log"
	"net"
)

// GetTCPListener gets a TCP listener choosing the first available port
func GetTCPListener(port string) (*net.TCPListener, string) {
	address, _ := net.ResolveTCPAddr(TCPProtocol, port)
	listener, errListen := net.ListenTCP(TCPProtocol, address)
	if errListen != nil {
		log.Fatal(errListen.Error())
	}

	// If the param be :0, go will pick automatically the port
	// This line gets the port chosen
	port = fmt.Sprint(listener.Addr().(*net.TCPAddr).Port)

	return listener, buildAddress(port)
}

// GetTCPDialer gets a TCP dialer to communicate with a TCP socket
func GetTCPDialer(address string) *net.TCPConn {
	tcpAddress, _ := net.ResolveTCPAddr(TCPProtocol, address)
	dialer, _ := net.DialTCP(TCPProtocol, nil, tcpAddress)
	return dialer
}

// buildAddress builds a network address compose by [host_ip]:[port]
func buildAddress(port string) string {
	return getLocalIP() + ":" + port
}

// getLocalIP gets the host ip address
func getLocalIP() string {
	addresses, err := net.InterfaceAddrs()
	var hostIP string
	if err != nil {
		hostIP = ""
	}

	index := 0
	for hostIP == "" && index < len(addresses) {
		address := addresses[index]

		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				hostIP = ipnet.IP.String()
			}
		}

		index++
	}

	return hostIP
}
