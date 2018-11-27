package internal

// Constants
const (
	MsgRunningServicesRegistration = "Running remote service registration at \"%s\""
	MsgRunningServicesLookup       = "Running remote service look-up at \"%s\""
	MsgRunningServicesInvoke       = "Running remote service invoke at \"%s\""

	MsgRemoteServiceConnectionIsOff = "Remote service \"%s:%s\" is off-line. Error \"%s\""
	MsgNamingServiceConnectionIsOff = "Naming service \"%s\" is off-line. Error \"%s\""

	MsgRegisteringService               = "Registering remote services \"%s\" from \"%s\""
	MsgUnregisterServiceAtNamingService = "Unregistering remote service \"%s:%s\""
	MsgRemoteServicesRegistrationStatus = "Registration status: #%d service(s): %s. Addresses: %s"
	MsgLookupRemoteService              = "Looking for remote service \"%s\""
	MsgFoundRemoteService               = "Found remote service \"%s\" from \"%s\""

	MsgClientInvokeRequest     = "Client invoke requests from \"%s\""
	MsgClientLookupRequest     = "Client look-up requests from \"%s\""
	MsgMethodNotFoundInService = "\"%s\" was not found in \"%s\""
	MsgServiceNotFound         = "\"%s\" was not found"
	MsgInvokingRemoteService   = "Invoking %s.%s(%s)"

	ArgNameLookupServerAddress       = "lookup_server_address"
	ArgLookupServerAddressMsg        = "The service look-up server address"
	ArgNameRegistrationServerAddress = "registration_server_address"
	ArgRegistrationServerAddressMsg  = "The service registration server address"
	ArgNameLookupServerPort          = "lookup_server_port"
	ArgLookupServerPortMsg           = "The service look-up server port"
	ArgNameRegistrationServerPort    = "registration_server_port"
	ArgRegistrationServerPortMsg     = "The service registration server port"
	ArgNameServerPort                = "port"
	ArgServerPortMsg                 = "The port where the server will run. (e.g: :8000)"
	ArgServerPortDefault             = ":0"
	ArgNameEncryptKey                = "encrypt_key"
	ArgEncryptKeyMsg                 = "The md5 encryptKey to encrypt/decrypt requests and responses"
	ArgEncryptKeyDefault             = "3a878ccc079a675df83041f1b695df6f"
)
