package network

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/marlom-jobsom/gorpc/internal/network/request"
	"github.com/marlom-jobsom/gorpc/internal/network/response"
)

// NewMarshaller build new instance of marshaller
func NewMarshaller(encryptKey string) *Marshaller {
	return &Marshaller{
		cipher: NewCipher(encryptKey),
	}
}

// Marshaller handles request / response data serialization and deserialization
type Marshaller struct {
	cipher *Cipher
}

// MarshalNamingServiceRegistration serializes a request
func (m *Marshaller) MarshalNamingServiceRegistration(namingServiceRegistration *request.NamingServiceRegistration) *request.NamingServiceRegistrationBytes {
	clientInvokeRequestBytes, err := json.Marshal(namingServiceRegistration)
	if err != nil {
		log.Fatal(err.Error())
	}

	clientInvokeRequestBytes = m.cipher.Encrypt(clientInvokeRequestBytes)
	clientInvokeRequestBytesSize := len(clientInvokeRequestBytes)
	return request.NewNamingServiceRegistrationBytes(
		clientInvokeRequestBytesSize, clientInvokeRequestBytes,
	)
}

// UnmarshalNamingServiceRegistration deserializes a request
func (m *Marshaller) UnmarshalNamingServiceRegistration(connection net.Conn) *request.NamingServiceRegistration {
	var registrationRequest request.NamingServiceRegistration
	var registrationRequestByte request.NamingServiceRegistrationBytes

	json.NewDecoder(connection).Decode(&registrationRequestByte)
	json.Unmarshal(m.cipher.Decrypt(registrationRequestByte.Bytes), &registrationRequest)

	return &registrationRequest
}

// MarshallLookupResponse serializes the lookup naming service server response
func (m *Marshaller) MarshallLookupResponse(address string) []byte {
	addressBytes, err := json.Marshal(address)
	if err != nil {
		log.Fatal(err.Error())
	}

	return m.cipher.Encrypt(addressBytes)
}

// UnmarshallLookupResponse deserializes the lookup naming service server response
func (m *Marshaller) UnmarshallLookupResponse(httpResponse *http.Response) string {
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var content string
	json.Unmarshal(m.cipher.Decrypt(body), &content)

	return content
}

// MarshalClientInvokeRequest serializes a client invoke request
func (m *Marshaller) MarshalClientInvokeRequest(clientInvokeRequest *request.ClientInvoke) *bytes.Buffer {
	requestBytes, err := json.Marshal(clientInvokeRequest)
	if err != nil {
		log.Fatal(err.Error())
	}

	return bytes.NewBuffer(m.cipher.Encrypt(requestBytes))
}

// UnmarshalClientInvokeRequest deserializes a client invoke request
func (m *Marshaller) UnmarshalClientInvokeRequest(htttpRequest *http.Request) *request.ClientInvoke {
	body, err := ioutil.ReadAll(htttpRequest.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var invokeRequest *request.ClientInvoke
	json.Unmarshal(m.cipher.Decrypt(body), &invokeRequest)

	return invokeRequest
}

// MarshalClientResponse serializes a response
func (m *Marshaller) MarshalClientResponse(response interface{}) []byte {
	responseByte, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err.Error())
	}

	return m.cipher.Encrypt(responseByte)
}

// UnmarshalClientResponse deserializes a response
func (m *Marshaller) UnmarshalClientResponse(httpResponse *http.Response) response.Response {
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var contentResponse response.Response
	json.Unmarshal(m.cipher.Decrypt(body), &contentResponse.Content)

	return contentResponse
}
