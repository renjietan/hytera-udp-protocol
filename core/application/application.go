package application

import (
	"github.com/renjietan/hytera-udp-protocol/core"
	"github.com/renjietan/hytera-udp-protocol/types"
)

func TempApplication(info ApplicationInfo) (types.UdpRequest, error) {
	res, err := core.TempBase(info.UserId, 0x01)
	if err != nil {
		return nil, err
	}
	res = append(res, types.Item{
		Name: "Payload",
		Value: types.UdpRequest{{
			Name:  "FuncPoint",
			Value: info.FuncPoint,
			Size:  2,
		}, {
			Name:  "ChannelNo",
			Value: info.ChannelNo,
			Size:  1,
		}, {
			Name:  "EventID",
			Value: info.EventID,
			Size:  2,
		}, {
			Name:  "EventRN",
			Value: info.EventRN,
			Size:  2,
		}, {
			Name: "AtomOpt",
			Value: types.UdpRequest{{
				Name:  "OptType",
				Value: info.AtomOpt.OptType,
				Size:  1,
			}, {
				Name:  "OptCode",
				Value: info.AtomOpt.OptCode,
				Size:  2,
			}, {
				Name:  "CallType",
				Value: info.AtomOpt.CallType,
				Size:  1,
			}, {
				Name:  "ParaSize",
				Value: nil,
				Size:  2,
			}, {
				Name:  "CallPara",
				Value: nil,
				Size:  0,
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}

type ApplicationInfo struct {
	UserId    int
	FuncPoint int
	ChannelNo int
	EventID   int
	EventRN   int
	AtomOpt   struct {
		OptType  int
		OptCode  int
		CallType int
		ParaSize int
	}
}
