package server

import (
	"log"
	"net/http"
	"reflect"

	"github.com/marlom-jobsom/gorpc/internal"
	"github.com/marlom-jobsom/gorpc/internal/network"
	"github.com/marlom-jobsom/gorpc/internal/network/request"
	"github.com/marlom-jobsom/gorpc/internal/services"
)

// NewInvoker builds a new instance of Invoker
func NewInvoker(encryptKey string) *Invoker {
	return &Invoker{
		RemoteService: services.NewRemoteService(encryptKey),
		marshaller:    network.NewMarshaller(encryptKey),
	}
}

// Invoker responsible for run method request by client
type Invoker struct {
	RemoteService *services.RemoteService
	marshaller    *network.Marshaller
}

// Invoke runs method requested
func (i *Invoker) Invoke(request *http.Request) []byte {
	clientInvoke := i.marshaller.UnmarshalClientInvokeRequest(request)
	output := i.invoke(clientInvoke)
	return i.marshaller.MarshalClientResponse(output)
}

// invoke runs method requested
func (i *Invoker) invoke(clientInvoke *request.ClientInvoke) interface{} {
	log.Printf(internal.MsgInvokingRemoteService,
		clientInvoke.ServiceName, clientInvoke.MethodName, clientInvoke.Arguments,
	)
	service := i.getService(clientInvoke.ServiceName)
	method := i.getMethod(service, clientInvoke.MethodName)
	arguments := i.getArguments(method, clientInvoke.Arguments)
	outputs := method.Call(arguments)
	return i.getMethodReturn(outputs)
}

// getService gets the service requested from service name
func (i *Invoker) getService(serviceName string) reflect.Value {
	serviceValue := reflect.ValueOf(i.RemoteService.GetService(serviceName))
	if !serviceValue.IsValid() {
		log.Fatalf(internal.MsgServiceNotFound, serviceName)
	}

	return serviceValue
}

// getMethod gets the method requested
func (i *Invoker) getMethod(service reflect.Value, methodName string) reflect.Value {
	methodType := service.MethodByName(methodName)

	if !methodType.IsValid() {
		log.Fatalf(internal.MsgMethodNotFoundInService, methodName, service.Type().String())
	}

	return methodType
}

// TODO: Add support to more types
// getArguments converts the arguments to their correct types for the method given
func (i *Invoker) getArguments(method reflect.Value, args []interface{}) []reflect.Value {
	argsValue := make([]reflect.Value, len(args))

	for index := range argsValue {
		arg := args[index]
		var newArg interface{}

		switch method.Type().In(index).Kind() {
		case reflect.Int:
			// Any numeric type from request is automatically converted to float64
			newArg = int(arg.(float64))
		case reflect.Float64:
			newArg = arg.(float64)
		case reflect.String:
			newArg = arg.(string)
		case reflect.Slice:
			newArg = arg
		}

		argsValue[index] = reflect.ValueOf(newArg)
	}

	return argsValue
}

// getMethodReturn converts the methods returns to their correct types
func (i *Invoker) getMethodReturn(outputs []reflect.Value) interface{} {
	outputsInterface := make([]interface{}, len(outputs))

	for index := range outputsInterface {
		outputsInterface[index] = outputs[index].Interface()
	}

	return outputsInterface
}
