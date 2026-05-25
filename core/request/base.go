package request

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core/request/types"
)

func TempBase(userId int, SAP int) (types.UdpRequestBytesCode, error) {
	if userId > 255 || userId < 0 {
		return nil, errors.New("userId 应在 0 - 255之间")
	}
	return types.UdpRequestBytesCode{{
		Name:  "SrcID", // 源模块ID。
		Value: 0xEE,    // 该字段取固定值0xEE
		Size:  1,
	}, {
		Name:  "DstID", // 目的模块ID。该字段取固定值0xEE
		Value: 0xEE,
		Size:  1,
	}, {
		Name:  "Length", // 长度。表示了该字段之后若干字段的字节数总和。
		Value: nil,
		Size:  2,
	}, {
		Name:  "CRC", // 校验码。表示了该字段之后所有数据的CRC校验值，
		Value: 0x00,  // 采用标准CRC16查表算法；填0，表示该消息无需校验
		Size:  2,
	}, {
		Name:  "Version", // 协议版本号。标记消息所使用的协议版本，从修订记录获取当前协议版本号；
		Value: 0x00,      // 例如，V0.1在该字段中表示成0x0
		Size:  1,
	}, {
		Name:  "UserID", // 客户端用户ID。为支持多用户连接电台，用户通过鉴权业务获得一个唯一用户ID，后续业务报文携带该用户ID，0x00为无效值
		Value: userId,
		Size:  1,
	}, {
		Name:  "SAP", // 业务接入点
		Value: SAP,
		Size:  1,
	}}, nil
}
