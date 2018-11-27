package services

import (
	"log"
	"net"

	"github.com/marlom-jobsom/gorpc/internal"
	"github.com/marlom-jobsom/gorpc/internal/network"
)

// NewNamingService builds a new instance of NamingService
func NewNamingService(encryptKey string) *NamingService {
	return &NamingService{
		remoteServicesEntries: make(map[string][]*Entry),
		marshaller:            network.NewMarshaller(encryptKey),
	}
}

// NamingService is a naming service who holds the urls to all services available for client
type NamingService struct {
	remoteServicesEntries map[string][]*Entry
	marshaller            *network.Marshaller
}

// RegisterServices registers new services that are available for client
func (n *NamingService) RegisterServices(connection net.Conn) {
	registrationRequest := n.marshaller.UnmarshalNamingServiceRegistration(connection)
	remoteServiceEntries := NewRemoteServiceEntries(
		connection, registrationRequest.ServicesNames, registrationRequest.ServerAddress,
	)

	for index := range remoteServiceEntries {
		n.registerService(remoteServiceEntries[index])
	}
}

// LookupService gets the first address on the list of addresses for the naming service given
func (n *NamingService) LookupService(serviceName string) []byte {
	log.Printf(internal.MsgLookupRemoteService, serviceName)
	var address string
	var entry *Entry
	entries := n.remoteServicesEntries[serviceName]

	if len(entries) > 0 {
		entry = entries[0]
		address = entry.Address
		log.Printf(internal.MsgFoundRemoteService, entry.Name, entry.Address)
	}

	return n.marshaller.MarshallLookupResponse(address)
}

// registerService registers a new service that is available for clients
func (n *NamingService) registerService(entry *Entry) {
	n.makeEntriesList(entry.Name)

	if !n.addressExists(entry.Name, entry.Address) {
		log.Printf(internal.MsgRegisteringService, entry.Name, entry.Address)
		entries := n.remoteServicesEntries[entry.Name]
		n.remoteServicesEntries[entry.Name] = append(entries, entry)
		n.status()

		go n.watchRemoteService(entry)
	}
}

// watchRemoteService deletes the given Entry from the map
func (n *NamingService) watchRemoteService(entry *Entry) {
	WatchRemoteService(entry)

	log.Printf(internal.MsgUnregisterServiceAtNamingService, entry.Name, entry.Address)
	entries := n.remoteServicesEntries[entry.Name]

	if len(entries) == 1 {
		n.deleteEntriesList(entry.Name)
	} else {
		n.deleteEntryList(entry)
	}

	n.status()
}

// makeEntriesList ensures there is a list of entries for the service name given
func (n *NamingService) makeEntriesList(serviceName string) {
	if _, nameExists := n.remoteServicesEntries[serviceName]; !nameExists {
		n.remoteServicesEntries[serviceName] = make([]*Entry, 0)
	}
}

// addressExists checks if the url given already exists
func (n *NamingService) addressExists(serviceName string, address string) bool {
	entries := n.remoteServicesEntries[serviceName]
	lengthEntries := len(entries)
	exists := false
	index := 0

	for !exists && lengthEntries > 0 && index < lengthEntries {
		exists = entries[index].Address == address
		index++
	}

	return exists
}

// deleteEntriesList deletes a entire list of address for a given service name
func (n *NamingService) deleteEntriesList(serviceName string) {
	delete(n.remoteServicesEntries, serviceName)
}

// deleteEntryList deletes a single entry from the entries list given
func (n *NamingService) deleteEntryList(entry *Entry) {
	index := 0
	indexFound := false
	entries := n.remoteServicesEntries[entry.Name]
	lengthEntries := len(entries)

	for !indexFound && index < lengthEntries {
		if entry.Address == entries[index].Address {
			indexFound = true
			entries[index] = entries[lengthEntries-1]
			entries[lengthEntries-1] = nil // Avoid potential memory leak problem
			n.remoteServicesEntries[entry.Name] = entries[:lengthEntries-1]
		}
		index++
	}
}

// status logs the status of all remote services types and addresses for each remote service registered
func (n *NamingService) status() {
	names := make([]string, 0)
	mapAddress := make(map[string][]string)

	for key := range n.remoteServicesEntries {
		names = append(names, key)
		addresses := make([]string, 0)

		entries := n.remoteServicesEntries[key]
		for index := range entries {
			addresses = append(addresses, entries[index].Address)
		}

		mapAddress[key] = addresses
	}

	log.Printf(internal.MsgRemoteServicesRegistrationStatus, len(names), names, mapAddress)
}
