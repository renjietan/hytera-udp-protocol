package tools

import (
	"net"

	"github.com/renjietan/hytera-udp-protocol/types"
)

var Success = func(event string, data any, message string, address *net.UDPAddr) types.Envelope {
	return types.Envelope{
		Event:   event,
		Code:    0,
		Message: message,
		Data:    data,
		Address: address,
	}
}

var Error = func(event string, message string, address *net.UDPAddr) types.Envelope {
	return types.Envelope{
		Event:   event,
		Code:    1,
		Message: message,
		Data:    nil,
		Address: address,
	}
}

var Bytes2Struct = func(b []byte, v types.UdpResponse) error {
	return nil
}
