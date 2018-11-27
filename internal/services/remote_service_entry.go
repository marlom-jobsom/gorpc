package services

import (
	"net"
)

// NewRemoteServiceEntry ...
func NewRemoteServiceEntry(connection net.Conn, serviceName string, serviceAddress string) *Entry {
	return &Entry{
		Connection: connection,
		Name:       serviceName,
		Address:    serviceAddress,
	}
}

// NewRemoteServiceEntries ...
func NewRemoteServiceEntries(connection net.Conn, servicesNames []string, serverAddress string) []*Entry {
	entries := make([]*Entry, 0)
	for index := range servicesNames {
		entry := NewRemoteServiceEntry(connection, servicesNames[index], serverAddress)
		entries = append(entries, entry)
	}
	return entries
}

// Entry ..
type Entry struct {
	Connection net.Conn
	Name       string
	Address    string
}
