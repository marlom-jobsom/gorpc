package request

// NewClientInvoke build new instance of ClientInvoke
func NewClientInvoke(serviceName string, methodName string, arguments []interface{}) *ClientInvoke {
	return &ClientInvoke{ServiceName: serviceName, MethodName: methodName, Arguments: arguments}
}

// ClientInvoke wrappers a client request to execute a remote method
type ClientInvoke struct {
	ServiceName string
	MethodName  string
	Arguments   []interface{}
}
