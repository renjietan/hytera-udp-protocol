package types

import "net"

var Success = func(event string, data any, message string, address *net.UDPAddr) Envelope {
	return Envelope{
		Event:   event,
		Code:    0,
		Message: message,
		Data:    data,
		Address: address,
	}
}

var Error = func(event string, message string, address *net.UDPAddr) Envelope {
	return Envelope{
		Event:   event,
		Code:    1,
		Message: message,
		Data:    nil,
		Address: address,
	}
}
