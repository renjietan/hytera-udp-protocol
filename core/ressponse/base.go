package response

import (
	types "github.com/renjietan/hytera-udp-protocol/core/ressponse/types"
)

func TempBase() types.UdpResponseBytesCode {
	return types.UdpResponseBytesCode{
		SrcID: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
		DstID: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
		Length: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  2,
		},
		CRC: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  2,
		},
		Version: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
		UserID: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
		SAP: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
	}
}

func TempPayload(optData any) types.Payload {
	return types.Payload{
		OptCode: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
		OptData: optData,
	}
}
