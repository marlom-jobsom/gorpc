package request

// NewNamingServiceRegistrationBytes build new instance of NamingServiceRegistrationBytes
func NewNamingServiceRegistrationBytes(size int, bytes []byte) *NamingServiceRegistrationBytes {
	return &NamingServiceRegistrationBytes{
		Size: size, Bytes: bytes,
	}
}

// NamingServiceRegistrationBytes wrappers the bytes of a request to register a service to be available for client
type NamingServiceRegistrationBytes struct {
	Size  int
	Bytes []byte
}
