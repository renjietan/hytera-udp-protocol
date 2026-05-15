package tools

import "net"

func ResolveUDPAddr(network, address string) (*net.UDPAddr, error) {
	return net.ResolveUDPAddr("udp", ":8080")
}
