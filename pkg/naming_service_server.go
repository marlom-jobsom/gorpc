package pkg

import (
	"log"
	"net/http"

	"github.com/marlom-jobsom/gorpc/internal"
	"github.com/marlom-jobsom/gorpc/internal/layers/infrastructure"
	"github.com/marlom-jobsom/gorpc/internal/network"
)

// NewNamingServiceServer builds a new instance of NamingServiceServer
func NewNamingServiceServer(lookupServerPort string, registrationServerPort string, encryptKey string) *NamingServiceServer {
	return &NamingServiceServer{
		requestHandler:         infrastructure.NewNamingServiceRequestHandler(encryptKey),
		lookupServerPort:       lookupServerPort,
		registrationServerPort: registrationServerPort,
	}
}

// NamingServiceServer handles the address for services available for clients
type NamingServiceServer struct {
	requestHandler         *infrastructure.NamingServiceRequestHandler
	lookupServerPort       string
	registrationServerPort string
}

// Run runs the naming service
func (n *NamingServiceServer) Run() {
	go n.runHTTPServerForServiceLookup()
	n.runSocketForServicesRegistration()
}

// runHTTPServerForServiceLookup runs a http server for remote services look-up
func (n *NamingServiceServer) runHTTPServerForServiceLookup() {
	listener, address := network.GetTCPListener(n.lookupServerPort)
	log.Printf(internal.MsgRunningServicesLookup, address)

	http.HandleFunc(network.LookupPath, n.requestHandler.HandleLookupServices)
	errServe := http.Serve(listener, nil)
	if errServe != nil {
		log.Fatal(errServe.Error())
	}
}

// runSocketForServicesRegistration runs a network socket for remote services registration
func (n *NamingServiceServer) runSocketForServicesRegistration() {
	listener, address := network.GetTCPListener(n.registrationServerPort)
	defer listener.Close()
	log.Printf(internal.MsgRunningServicesRegistration, address)

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println(err.Error())
		}

		go n.requestHandler.HandleRegistrationServices(connection)
	}
}
