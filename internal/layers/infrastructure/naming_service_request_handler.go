package infrastructure

import (
	"log"
	"net"
	"net/http"

	"github.com/marlom-jobsom/gorpc/internal"
	"github.com/marlom-jobsom/gorpc/internal/network"
	"github.com/marlom-jobsom/gorpc/internal/services"
)

// NewNamingServiceRequestHandler builds a new NamingServiceRequestHandler
func NewNamingServiceRequestHandler(encryptKey string) *NamingServiceRequestHandler {
	return &NamingServiceRequestHandler{
		namingService: *services.NewNamingService(encryptKey),
	}
}

// NamingServiceRequestHandler is responsible for handle remote service's registration requests
type NamingServiceRequestHandler struct {
	namingService services.NamingService
}

// HandleLookupServices handles client's look-up requests for available remote services
func (n *NamingServiceRequestHandler) HandleLookupServices(writer http.ResponseWriter, request *http.Request) {
	log.Printf(internal.MsgClientLookupRequest, request.RemoteAddr)
	serviceName := request.URL.EscapedPath()[len(network.LookupPath):]
	addressBytes := n.namingService.LookupService(serviceName)
	writer.Header().Set(network.HeaderContentTypeTag, network.HeaderApplicationJSONUTF8)
	writer.Write(addressBytes)
}

// HandleRegistrationServices handles remote services registration requests
func (n *NamingServiceRequestHandler) HandleRegistrationServices(connection net.Conn) {
	n.namingService.RegisterServices(connection)
}
