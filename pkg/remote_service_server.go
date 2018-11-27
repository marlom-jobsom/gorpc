package pkg

import (
	"log"
	"net"
	"net/http"

	"github.com/marlom-jobsom/gorpc/internal"
	"github.com/marlom-jobsom/gorpc/internal/layers/infrastructure"
	"github.com/marlom-jobsom/gorpc/internal/network"
)

// NewRemoteServiceServer builds a new instance of RemoteServiceServer
func NewRemoteServiceServer(port string, registrationServerAddress string, encryptKey string) *RemoteServiceServer {
	return &RemoteServiceServer{
		requestHandler: infrastructure.NewRemoteServiceRequestHandler(encryptKey),
		port:           port,
		registrationServerAddress: registrationServerAddress,
	}
}

// RemoteServiceServer holds services that can be invoke for clients
type RemoteServiceServer struct {
	requestHandler            *infrastructure.RemoteServiceRequestHandler
	port                      string
	registrationServerAddress string
}

// Run runs the remote service
func (r *RemoteServiceServer) Run() {
	listener, address := network.GetTCPListener(r.port)
	go r.runHTTPServerForServicesInvocation(listener, address)
	r.bindServicesInNamingService(address)
}

// RegisterServiceInNamingService adds a new service that will be available for clients
func (r *RemoteServiceServer) RegisterServiceInNamingService(name string, instance interface{}) {
	r.requestHandler.Invoker.RemoteService.RegisterService(name, instance)
}

// runHTTPServerForServicesInvocation brings up the http server that handles services invoke requests
func (r *RemoteServiceServer) runHTTPServerForServicesInvocation(listener net.Listener, address string) {
	log.Printf(internal.MsgRunningServicesInvoke, address)
	http.HandleFunc(network.InvokePath, r.requestHandler.HandleInvokeRequest)

	errServe := http.Serve(listener, nil)
	if errServe != nil {
		log.Fatal(errServe.Error())
	}
}

// bindServicesInNamingService binds services on naming service server
func (r *RemoteServiceServer) bindServicesInNamingService(address string) {
	r.requestHandler.Invoker.RemoteService.BindNamingService(address, r.registrationServerAddress)
}
