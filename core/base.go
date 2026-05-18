package core

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/options"
	"github.com/renjietan/hytera-udp-protocol/tools"
)

func TempBase(userId int) (options.UdpRequest, error) {
	if userId > 255 || userId < 0 {
		return nil, errors.New("userId 应在 0 - 255之间")
	}
	return options.UdpRequest{{
		Name:  "SrcID",
		Value: 0xEE,
		Size:  1,
	}, {
		Name:  "DstID",
		Value: 0xEE,
		Size:  1,
	}, {
		Name:  "Length",
		Value: nil,
		Size:  2,
	}, {
		Name:  "CRC",
		Value: 0x00,
		Size:  2,
	}, {
		Name:  "Version",
		Value: 0x00,
		Size:  1,
	}, {
		Name:  "UserID",
		Value: userId,
		Size:  1,
	}, {
		Name:  "SAP",
		Value: 0x01,
		Size:  1,
	}}, nil
}

func Struct2Bytes(params options.UdpRequest) []byte {
	refsField := tools.GetRecursiveField(params, []options.Item{})
	_, res := tools.Struct2Bytes(params, refsField, []byte{})
	return res
}
