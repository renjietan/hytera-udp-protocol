package application_extended

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// ReadFuncPointsListReq 读取功能点列表请求
var ReadFuncPointsListReq = func(SegmentIndex, userId int) ([]byte, error) {
	tempByte, err := TReadFuncPointsListReq(SegmentIndex, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the TReadFuncPointsListReq template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.Item{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

var TReadFuncPointsListReq = func(SegmentIndex, userId int) (types.UdpRequest, error) {
	res, err := core.TempBase(userId, 0x03)
	if err != nil {
		return nil, err
	}
	res = append(res, types.Item{
		Name: "Payload",
		Value: types.UdpRequest{{
			Name:  "OptCode",
			Value: 0x05,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequest{{
				Name:  "SegmentIndex", // 段编号，从1开始
				Value: SegmentIndex,
				Size:  1,
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}
