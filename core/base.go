package core

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

func TempBase(userId int, SAP int) (types.UdpRequest, error) {
	if userId > 255 || userId < 0 {
		return nil, errors.New("userId 应在 0 - 255之间")
	}
	return types.UdpRequest{{
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
		Value: SAP,
		Size:  1,
	}}, nil
}

func Struct2Bytes(params types.UdpRequest) []byte {
	refsField := tools.GetRecursiveField(params, []types.Item{})
	_, res := tools.Struct2Bytes(params, refsField, []byte{})
	return res
}
