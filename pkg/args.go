package pkg

import (
	"flag"

	"github.com/marlom-jobsom/gorpc/internal"
)

// GetClientArgs reads the parameters from CLI interface for client
func GetClientArgs() (string, string) {
	var lookupServerAddress string
	var encryptKey string

	flag.StringVar(
		&lookupServerAddress,
		internal.ArgNameLookupServerAddress,
		"",
		internal.ArgLookupServerAddressMsg,
	)
	flag.StringVar(
		&encryptKey,
		internal.ArgNameEncryptKey,
		internal.ArgEncryptKeyDefault,
		internal.ArgEncryptKeyMsg,
	)
	flag.Parse()

	return lookupServerAddress, encryptKey
}

// GetRemoteServiceArgs reads the parameters from CLI interface for remote service
func GetRemoteServiceArgs() (string, string, string) {
	var port string
	var registrationServerAddress string
	var encryptKey string

	flag.StringVar(
		&port,
		internal.ArgNameServerPort,
		internal.ArgServerPortDefault,
		internal.ArgServerPortMsg,
	)
	flag.StringVar(
		&registrationServerAddress,
		internal.ArgNameRegistrationServerAddress,
		"",
		internal.ArgRegistrationServerAddressMsg,
	)
	flag.StringVar(
		&encryptKey,
		internal.ArgNameEncryptKey,
		internal.ArgEncryptKeyDefault,
		internal.ArgEncryptKeyMsg,
	)
	flag.Parse()

	return port, registrationServerAddress, encryptKey
}

// GetNamingServiceArgs reads the parameters from CLI interface for naming service
func GetNamingServiceArgs() (string, string, string) {
	var lookupServerPort string
	var registrationServerPort string
	var encryptKey string

	flag.StringVar(
		&lookupServerPort,
		internal.ArgNameLookupServerPort,
		internal.ArgServerPortDefault,
		internal.ArgLookupServerPortMsg,
	)
	flag.StringVar(
		&registrationServerPort,
		internal.ArgNameRegistrationServerPort,
		internal.ArgServerPortDefault,
		internal.ArgRegistrationServerPortMsg,
	)
	flag.StringVar(
		&encryptKey,
		internal.ArgNameEncryptKey,
		internal.ArgEncryptKeyDefault,
		internal.ArgEncryptKeyMsg,
	)
	flag.Parse()

	return lookupServerPort, registrationServerPort, encryptKey
}
