package options

import "net"

type Envelope struct {
	Event   string
	Code    int
	Message string
	Data    interface{}
	Address *net.UDPAddr
}
