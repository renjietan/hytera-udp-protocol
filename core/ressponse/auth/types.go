package core_response_auth

import (
	response "github.com/renjietan/hytera-udp-protocol/core/ressponse"
	"github.com/renjietan/hytera-udp-protocol/core/ressponse/types"
)

// LoginAck 登录-0x81 0x85
type LoginAck struct {
	// 0x00    成功
	// 0x01    当前登录用户数量达到上限
	// 0x02    用户已经登录
	// 0x03    错误，用户名非法
	Status types.UdpResponseByteCodeItem // 1 Byte
	// 由电台动态生成，标识本次连接的唯一性，0为无效值
	UserID types.UdpResponseByteCodeItem // 1 Byte
}

func NewLoginAck() types.Payload {
	res := LoginAck{
		Status: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
		UserID: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
	}
	return response.TempPayload(res)
}

// LogoutAck 登出-0x82
type LogoutAck struct {
	// 0x00    成功。
	// 0x01    无此用户。
	// 0x02    用户名与用户ID不匹配
	Status types.UdpResponseByteCodeItem // 1 Byte
}

func NewLogoutAck() types.Payload {
	res := LogoutAck{
		Status: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
	}
	return response.TempPayload(res)
}

// PingAck 心跳-0x83
type PingAck struct {
	// 0x00    正常。
	// 0x01    运行错误。
	Status types.UdpResponseByteCodeItem // 1 Byte
}

func NewPingAck() types.Payload {
	res := PingAck{
		Status: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
	}
	return response.TempPayload(res)
}

// PasswordAck 密码-0x84
type PasswordAck struct {
	// 0x01    密码校验成功
	// 0x02    密码校验失败
	Status types.UdpResponseByteCodeItem // 1 Byte
	// 0x01    普通权限
	// 0x02    高级权限
	// 0x03    超级权限
	// 0x10~0xFF 项目自定义权限
	// 权限表由项目自行提供。
	PermissionLevel types.UdpResponseByteCodeItem // 1 Byte
}

func NewPasswordAck() types.Payload {
	res := PasswordAck{
		Status: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
		PermissionLevel: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
	}
	return response.TempPayload(res)
}

// KickOutSubAck 电台向用户发送踢出用户订阅的应答-0x86
type KickOutSubAck struct {
	// 0x00    成功
	// 0x01    失败
	Status types.UdpResponseByteCodeItem // 1 Byte
}

func NewKickOutSubAck() types.Payload {
	res := KickOutSubAck{
		Status: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
	}
	return response.TempPayload(res)
}

// KickOutUserAck 电台向用户发送踢出用户请求的应答-0x87
type KickOutUserAck struct {
	// 0x00    成功
	// 0x01    权限不足
	// 0x02    用户未登录
	Status types.UdpResponseByteCodeItem // 1 Byte
}

func NewKickOutUserAck() types.Payload {
	res := KickOutUserAck{
		Status: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
	}
	return response.TempPayload(res)
}

// KickOutInfoNty 电台向用户发送踢出用户信息通知-0x88
type KickOutInfoNty struct {
	// 用户名的字的个数，最大40个字。
	Size types.UdpResponseByteCodeItem // 1 Byte
	// UTF-16BE编码。最大长度40个字。用户名
	UserName types.UdpResponseByteCodeItem // n Byte
	// 离线时间：
	// 0    永久
	// >0   秒数
	OfflineTime types.UdpResponseByteCodeItem // 4 Byte
}

func NewKickOutInfoNty() types.Payload {
	res := KickOutInfoNty{
		Size: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
			Bind: types.UdpResponseBindStruct{
				Path:      "Payload.OptData.UserName",
				FieldName: types.UdpResponseBindStructSize,
				//Callback: func(params ...any) {
				//	fmt.Println("NewKickOutInfoNty Func===============", params)
				//},
				Value:     nil,
				ValueType: types.UdpResponseBindStructInt,
			}, // 下面 UserName 字段
		},
		UserName: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  0,
		},
		OfflineTime: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
	}
	return response.TempPayload(res)
}

// KickOutInfoAck 电台向用户发送踢出用户信息通知的应答-0x08
type KickOutInfoAck struct {
	// 0: 正常    1: 故障
	Status types.UdpResponseByteCodeItem // 1 Byte
}

func NewKickOutInfoAck() types.Payload {
	res := KickOutInfoAck{
		Status: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
	}
	return response.TempPayload(res)
}

// LoginInfoAck 电台收到获取登录信息请求之后，返回的应答-0x89
type LoginInfoAck struct {
	// 查询用户登录信息结果：
	// 0x00    成功
	// 0x01    未知用户
	// 0x02    权限不足
	Result types.UdpResponseByteCodeItem // 1 Byte
	// 心跳检测间隔。单位：毫秒。最小值：100
	SuperviseInterval types.UdpResponseByteCodeItem // 4 Byte
	// 心跳检测次数。
	// > 2        0、1、2为无效值；
	// 0xffff    不启用心跳功能。
	SuperviseCnt types.UdpResponseByteCodeItem // 2 Byte
	// 用户登录时间，见附录中“时间戳定义”。
	LoginTime types.UdpResponseByteCodeItem // n Byte
	// 最近一次交互时间，见附录中“时间戳定义”
	RecentTime types.UdpResponseByteCodeItem // n Byte
	// 设备适配器所属通道号，0 代表不区分；从1开始，表示通道一、通道二等等。
	ChannelNo types.UdpResponseByteCodeItem // 1 Byte
	// 设备适配器类型：
	// 0x00    无效
	// 0x01    以太网
	// 0x02    串口
	// 0x03    被覆线
	AdapterType types.UdpResponseByteCodeItem // 1 Byte
	// 设备适配器类型：
	// 0x00    无效
	// 0x01    以太网
	// 0x02    串口
	// 0x03    被覆线
	Ipv4Addr types.UdpResponseByteCodeItem // 4 Byte
	//设备适配器IPv4掩码。AdapterType为0x01时有效
	Ipv4Mask types.UdpResponseByteCodeItem // 4 Byte
	// 设备适配器名称字的个数，最大255个字
	Size types.UdpResponseByteCodeItem // 1 Byte
	// 设备适配器名称: UTF-16BE编码
	AdapterName types.UdpResponseByteCodeItem // n Byte
}

func NewLoginInfoAck() types.Payload {
	res := LoginInfoAck{
		Result: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
		SuperviseInterval: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  4,
		},
		SuperviseCnt: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  2,
		},
		LoginTime: types.UdpResponseByteCodeItem{
			Value: []byte{},
		},
		RecentTime: types.UdpResponseByteCodeItem{
			Value: []byte{},
		},
		ChannelNo: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
		AdapterType: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
		Ipv4Addr: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  4,
		},
		Ipv4Mask: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  4,
		},
		Size: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
		AdapterName: types.UdpResponseByteCodeItem{
			Value: []byte{},
		},
	}
	return response.TempPayload(res)
}

// AdapterInfoAck 电台收到获取所有适配器信息请求之后，返回的应答-0x0A
type AdapterInfoAck struct {
	// 查询用户登录信息结果：
	// 0x00    成功
	// 0x01    权限不足
	Result types.UdpResponseByteCodeItem // 1 Byte
	// 设备适配器链表元素数量
	ListCount types.UdpResponseByteCodeItem // 2 Bytes
	// 适配器列表
	Channels []AdapterInfoItem
}

func NewAdapterInfoAck() types.Payload {
	res := AdapterInfoAck{
		Result: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
		// 设备适配器链表元素数量
		ListCount: types.UdpResponseByteCodeItem{
			Value: []byte{},
			Size:  1,
		},
		// 根据 ListCount 计算
		Channels: make([]AdapterInfoItem, 0),
	}
	return response.TempPayload(res)
}

type AdapterInfoItem struct {
	// 设备适配器所属通道号, 0 代表不区分；从1开始，表示通道一、通道二等等。
	ChannelNo types.UdpResponseByteCodeItem // 1 Byte
	// 设备适配器类型：
	// 0x00    无效
	// 0x01    以太网
	// 0x02    串口
	// 0x03    被覆线
	AdapterType types.UdpResponseByteCodeItem // 1 Byte
	// 设备适配器IPv4地址。AdapterType为0x01时有效
	Ipv4Addr types.UdpResponseByteCodeItem // 4 Byte
	// 设备适配器IPv4掩码。AdapterType为0x01时有效
	Ipv4Mask types.UdpResponseByteCodeItem // 4 Byte
	// 设备适配器名称字的个数，最大255个字
	Size types.UdpResponseByteCodeItem // 1 Byte
	// 设备适配器名称: UTF-16BE编码
	AdapterName types.UdpResponseByteCodeItem // n Byte
}
