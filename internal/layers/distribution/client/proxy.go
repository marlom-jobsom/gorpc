package client

import (
	"github.com/marlom-jobsom/gorpc/internal/network/response"
)

// NewProxy builds a new instance of Proxy
func NewProxy(lookupServerAddress string, encryptKey string, serviceName string) *Proxy {
	return &Proxy{
		requestor:   NewRequestor(lookupServerAddress, encryptKey),
		serviceName: serviceName,
	}
}

// Proxy communicate with the remote service
type Proxy struct {
	requestor   *Requestor
	serviceName string
}

// Invoke runs method on the remote service
func (p *Proxy) Invoke(methodName string, arguments ...interface{}) response.Response {
	return p.requestor.Invoke(p.serviceName, methodName, arguments)
}
