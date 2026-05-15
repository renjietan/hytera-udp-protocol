package core

import (
	"fmt"

	"github.com/renjietan/hytera-udp-protocol/options"
)

var Login = options.UdpRequest{
	"SrcID":   0xEE,
	"DstID":   0xEE,
	"Length":  nil,
	"CRC":     0x00,
	"Version": 0x00,
	"UserID":  0x00,
	"SAP":     0x01,
	"Payload": options.UdpRequest{
		"OptCode": 0x01,
		"OptData": options.UdpRequest{
			"Size":     nil,
			"UserName": 5,
		},
	},
}

func Map2Bt(params options.UdpRequest) {
	for k, v := range params {
		fmt.Println(k, v)
	}
}
