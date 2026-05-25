package core_request_auth

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core/request"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// AdapterInfo
// 设备适配器类型：用户可以获取设备的所有适配器信息
//   - 0x00    不限
//   - 0x01    以太网
//   - 0x02    串口
//   - 0x03    被覆线
var AdapterInfoReq = func(AdapterType int, userId int) ([]byte, error) {
	tempByte, err := TAdapterInfoReq(AdapterType, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the AdapterInfoReq template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.UdpRequestByteCodeItem{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

var TAdapterInfoReq = func(AdapterType int, userId int) (types.UdpRequestBytesCode, error) {
	res, err := request.TempBase(userId, 0x01)
	if err != nil {
		return nil, err
	}
	res = append(res, types.UdpRequestByteCodeItem{
		Name: "Payload",
		Value: types.UdpRequestBytesCode{{
			Name:  "OptCode", // 操作码
			Value: 0x0a,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequestBytesCode{{
				// 设备适配器类型：
				// 0x00    不限
				// 0x01    以太网
				// 0x02    串口
				// 0x03    被覆线
				Name:  "AdapterType",
				Value: AdapterType,
				Size:  1,
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}
