package services

import (
	"log"
	"net"

	"github.com/marlom-jobsom/gorpc/internal"
)

// WatchRemoteService keeps monitoring the given Entry connection
func WatchRemoteService(entry *Entry) {
	err := watch(entry.Connection)
	log.Printf(internal.MsgRemoteServiceConnectionIsOff,
		entry.Name, entry.Address, err,
	)
}

// WatchNamingService keeps monitoring the given naming service connection
func WatchNamingService(connection net.Conn) {
	err := watch(connection)
	log.Printf(internal.MsgNamingServiceConnectionIsOff,
		connection.RemoteAddr().String(), err,
	)
}

// watch keeps monitoring a given connection
func watch(connection net.Conn) error {
	var err error

	for {
		_, err = connection.Read(make([]byte, 1))
		if err != nil {
			connection.Close()
			break
		}
	}

	return err
}
