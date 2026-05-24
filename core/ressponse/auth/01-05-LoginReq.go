package core_response_auth

import (
	"github.com/renjietan/hytera-udp-protocol/core/ressponse"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// LoginRes 登录
var LoginRes = func() types.UdpRequestBytesCode {
	res := ressponse.TempBase()
	return res
}
