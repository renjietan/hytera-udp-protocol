package core

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// AdapterInfo
// 设备适配器类型：用户可以获取设备的所有适配器信息
//   - 0x00    不限
//   - 0x01    以太网
//   - 0x02    串口
//   - 0x03    被覆线
var AdapterInfo = func(AdapterType int, userId int) ([]byte, error) {
	tempByte, err := TempAdapterInfo(AdapterType, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the Adapter Info template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.Item{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

var TempAdapterInfo = func(AdapterType int, userId int) (types.UdpRequest, error) {
	res, err := TempBase(userId, 0x01)
	if err != nil {
		return nil, err
	}
	res = append(res, types.Item{
		Name: "Payload",
		Value: types.UdpRequest{{
			Name:  "OptCode",
			Value: 0x0a,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequest{{
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
