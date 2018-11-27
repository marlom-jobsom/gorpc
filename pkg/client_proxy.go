package pkg

import (
	"github.com/marlom-jobsom/gorpc/internal/layers/distribution/client"
)

// NewClientProxy builds a new instance of clientProxy
func NewClientProxy(lookupServerAddress string, encryptKey string, serviceName string) *ClientProxy {
	return &ClientProxy{Proxy: client.NewProxy(lookupServerAddress, encryptKey, serviceName)}
}

// ClientProxy communicate with the remote service
type ClientProxy struct {
	*client.Proxy
}
