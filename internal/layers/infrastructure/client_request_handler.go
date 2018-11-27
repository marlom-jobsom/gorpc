package infrastructure

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/marlom-jobsom/gorpc/internal/network"
)

// NewClientRequestHandler builds a new instance of ClientRequestHandler
func NewClientRequestHandler() *ClientRequestHandler {
	return &ClientRequestHandler{
		Client: &http.Client{Timeout: time.Second * 10},
	}
}

// ClientRequestHandler is responsible for sending request to remote service
type ClientRequestHandler struct {
	*http.Client
}

// Lookup looks for a remote service address for the naming service given
func (r *ClientRequestHandler) Lookup(lookupServerAddress string, serviceName string) *http.Response {
	response, err := r.Get(
		network.HTTPProtocol + lookupServerAddress + network.LookupPath + serviceName,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	return response
}

// Send sends a invoke request to remote service
func (r *ClientRequestHandler) Send(remoteServiceAddress string, request *bytes.Buffer) *http.Response {
	response, err := r.Post(
		network.HTTPProtocol+remoteServiceAddress+network.InvokePath,
		network.RequestTypeApplicationJSON,
		request,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	return response
}
