package tools

import (
	"errors"
	"fmt"
	"net"
)

func GenerateAddress(RHost string, RPort int) (*net.UDPAddr, error) {
	// 校验: 远程IP
	address := fmt.Sprintf("%s:%v", RHost, RPort)
	rAddress, rErr := net.ResolveUDPAddr("udp", address)
	if rErr != nil {
		return nil, errors.New("不是有效的地址")
	}
	return rAddress, nil
}
