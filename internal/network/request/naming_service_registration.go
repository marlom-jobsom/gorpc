package request

// NewNamingServiceRegistration build new instance of NamingServiceRegistration
func NewNamingServiceRegistration(serviceName []string, serverAddress string) *NamingServiceRegistration {
	return &NamingServiceRegistration{
		ServicesNames: serviceName, ServerAddress: serverAddress,
	}
}

// NamingServiceRegistration wrappers a request to register a service to be available for client
type NamingServiceRegistration struct {
	ServicesNames []string
	ServerAddress string
}
