package application

import (
	"errors"
	"math/rand"

	"github.com/renjietan/hytera-udp-protocol/core"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

// Application 应用功能业务 0x02
var Application = func(info types.ApplicationInfo, userId int) ([]byte, error) {
	tempByte, err := TempApplication(info, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the application template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.Item{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

func TempApplication(info types.ApplicationInfo, userId int) (types.UdpRequest, error) {
	res, err := core.TempBase(userId, 0x02)
	if err != nil {
		return nil, err
	}
	res = append(res, types.Item{
		Name: "Payload",
		Value: types.UdpRequest{{
			Name:  "FuncPoint",    // 功能点枚举，0为无效值
			Value: info.FuncPoint, // 例如: 见 2.5.3功能点 0x0008-波形管理功能点(功能点枚举，0为无效值)
			Size:  2,
		}, {
			Name:  "ChannelNo",    // 所属通道号，0 代表通用功能点；从1开始，表示通道一、通道二等等,
			Value: info.ChannelNo, // 例如:0x00
			Size:  1,
		}, {
			Name:  "EventID",        // 事务ID。用于实现并发业务:
			Value: rand.Intn(30000), // 例如: rand.Intn(30000) 0-30000随机数
			Size:  2,
		}, {
			Name:  "EventRN",        // 事务随机码。用于标识消息的唯一性
			Value: rand.Intn(30000), // 例如: rand.Intn(30000) 0-30000随机数
			Size:  2,
		}, {
			Name: "AtomOpt",
			Value: types.UdpRequest{{
				Name:  "OptType",            // 操作类型
				Value: info.AtomOpt.OptType, // 例如: 见 2.5.2.1操作类型; 0x10-波形操作(waveform)
				Size:  1,
			}, {
				Name:  "OptCode",            // 操作码。原子操作的唯一标识
				Value: info.AtomOpt.OptCode, // 例如: 见 2.5.2.10 (0x10)波形操作;
				Size:  2,
			}, {
				Name: "CallType", // 调用类型（高四位表示方向，低四位表示类型）。
				// 请求(REQ): 0x01
				// 订阅(SUB): 0x02-SUB
				// 通知(NTY): 0x03-NTY
				// 应答(ACK): 0x04-ACK
				Value: info.AtomOpt.CallType,
				Size:  1,
			}, {
				Name:  "ParaSize", // CallPara(调用参数)总长度。
				Value: nil,
				Size:  2,
			}, {
				Name:  "CallPara",            // 调用参数
				Value: info.AtomOpt.CallPara, // 例如: 见(2.5.2.10.7) (0x0007) 子网参数配置
				Size:  0,
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}
