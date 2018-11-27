package infrastructure

import (
	"log"
	"net/http"

	"github.com/marlom-jobsom/gorpc/internal"
	"github.com/marlom-jobsom/gorpc/internal/layers/distribution/server"
	"github.com/marlom-jobsom/gorpc/internal/network"
)

// NewRemoteServiceRequestHandler builds a new RemoteServiceRequestHandler
func NewRemoteServiceRequestHandler(encryptKey string) *RemoteServiceRequestHandler {
	return &RemoteServiceRequestHandler{
		Invoker: server.NewInvoker(encryptKey),
	}
}

// RemoteServiceRequestHandler is responsible for handle client's invocation requests
type RemoteServiceRequestHandler struct {
	Invoker *server.Invoker
}

// HandleInvokeRequest handles client's requests
func (r *RemoteServiceRequestHandler) HandleInvokeRequest(writer http.ResponseWriter, request *http.Request) {
	log.Printf(internal.MsgClientInvokeRequest, request.RemoteAddr)
	output := r.Invoker.Invoke(request)
	writer.Header().Set(network.HeaderContentTypeTag, network.HeaderApplicationJSONUTF8)
	writer.Write(output)
}
