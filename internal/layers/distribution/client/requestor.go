package client

import (
	"bytes"
	"net/http"
	"time"

	"github.com/marlom-jobsom/gorpc/internal/layers/infrastructure"
	"github.com/marlom-jobsom/gorpc/internal/network"
	"github.com/marlom-jobsom/gorpc/internal/network/request"
	"github.com/marlom-jobsom/gorpc/internal/network/response"
)

// NewRequestor builds a new instance of Requestor
func NewRequestor(lookupServerAddress string, encryptKey string) *Requestor {
	return &Requestor{
		requestHandler:      infrastructure.NewClientRequestHandler(),
		marshaller:          network.NewMarshaller(encryptKey),
		lookupServerAddress: lookupServerAddress,
	}
}

// Requestor is responsible access the remote service
type Requestor struct {
	requestHandler      *infrastructure.ClientRequestHandler
	marshaller          *network.Marshaller
	lookupServerAddress string
}

// Invoke invokes a method on a remote service
func (r *Requestor) Invoke(serviceName string, methodName string, arguments []interface{}) response.Response {
	remoteServiceAddress := r.lookup(serviceName)
	requestData := r.marshal(serviceName, methodName, arguments)
	serverResponse, duration := r.send(remoteServiceAddress, requestData)
	return r.unmarshal(serverResponse, duration)
}

// lookup looks for a remote service address for the naming service given
func (r *Requestor) lookup(serviceName string) string {
	lookUpResponse := r.requestHandler.Lookup(r.lookupServerAddress, serviceName)
	return r.marshaller.UnmarshallLookupResponse(lookUpResponse)
}

/// marshal serializes an invoke request
func (r *Requestor) marshal(serviceName string, methodName string, arguments []interface{}) *bytes.Buffer {
	clientInvokeRequest := request.NewClientInvoke(serviceName, methodName, arguments)
	return r.marshaller.MarshalClientInvokeRequest(clientInvokeRequest)
}

// send sends a invoke request to remote service
func (r *Requestor) send(remoteServiceAddress string, requestData *bytes.Buffer) (*http.Response, time.Duration) {
	now := time.Now()
	serverResponse := r.requestHandler.Send(remoteServiceAddress, requestData)
	elapsed := time.Since(now)
	return serverResponse, elapsed
}

// unmarshal deserializes an http response
func (r *Requestor) unmarshal(serverResponse *http.Response, duration time.Duration) response.Response {
	clientResponse := r.marshaller.UnmarshalClientResponse(serverResponse)
	clientResponse.Duration = duration
	return clientResponse
}
