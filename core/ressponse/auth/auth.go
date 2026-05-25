package core_response_auth

import (
	"github.com/renjietan/hytera-udp-protocol/core/ressponse"
	types "github.com/renjietan/hytera-udp-protocol/core/ressponse/types"
)

// LoginAckRes 登录-0x81 0x85
var LoginAckRes = func() types.UdpResponseBytesCode {
	res := response.TempBase()
	payload := NewLoginAck()
	res.Payload = payload
	return res
}

// LogoutAckRes 登出-0x82
var LogoutAckRes = func() types.UdpResponseBytesCode {
	res := response.TempBase()
	payload := NewLogoutAck()
	res.Payload = payload
	return res
}

// PingAckRes 心跳-0x83
var PingAckRes = func() types.UdpResponseBytesCode {
	res := response.TempBase()
	payload := NewPingAck()
	res.Payload = payload
	return res
}

// PasswordAckRes 密码-0x84
var PasswordAckRes = func() types.UdpResponseBytesCode {
	res := response.TempBase()
	payload := NewPasswordAck()
	res.Payload = payload
	return res
}

// KickOutSubAckRes 电台向用户发送踢出用户订阅的应答-0x86
var KickOutSubAckRes = func() types.UdpResponseBytesCode {
	res := response.TempBase()
	payload := NewKickOutSubAck()
	res.Payload = payload
	return res
}

// KickOutUserAckRes 电台向用户发送踢出用户请求的应答-0x87
var KickOutUserAckRes = func() types.UdpResponseBytesCode {
	res := response.TempBase()
	payload := NewKickOutUserAck()
	res.Payload = payload
	return res
}

// KickOutUserNtyRes 电台向用户发送踢出用户信息通知-0x88
var KickOutUserNtyRes = func() types.UdpResponseBytesCode {
	res := response.TempBase()
	payload := NewKickOutInfoNty()
	res.Payload = payload
	return res
}

// KickOutInfoAckRes 电台向用户发送踢出用户信息通知的应答-0x08
var KickOutInfoAckRes = func() types.UdpResponseBytesCode {
	res := response.TempBase()
	payload := NewKickOutInfoAck()
	res.Payload = payload
	return res
}

// LoginInfoAckRes 电台收到获取登录信息请求之后，返回的应答-0x89
var LoginInfoAckRes = func() types.UdpResponseBytesCode {
	res := response.TempBase()
	payload := NewLoginInfoAck()
	res.Payload = payload
	return res
}

// NewAdapterInfoAckRes 电台收到获取所有适配器信息请求之后，返回的应答-0x0A
var NewAdapterInfoAckRes = func() types.UdpResponseBytesCode {
	res := response.TempBase()
	payload := NewAdapterInfoAck()
	res.Payload = payload
	return res
}
