package core_response_auth

import (
	"github.com/renjietan/hytera-udp-protocol/core/ressponse"
	types "github.com/renjietan/hytera-udp-protocol/types/reponse"
	types_response_auth "github.com/renjietan/hytera-udp-protocol/types/reponse/auth"
)

// LoginRes 登录
var LoginRes = func() types.UdpResponseBytesCode {
	res := response.TempBase()
	optData := types_response_auth.NewLoginAck()
	payload := types.Payload{
		OptCode: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
		OptData: optData,
	}
	res.Payload = payload
	return res
}
