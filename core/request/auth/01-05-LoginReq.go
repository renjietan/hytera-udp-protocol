package core_request_auth

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/core/request"
	"github.com/renjietan/hytera-udp-protocol/core/request/types"
	"github.com/renjietan/hytera-udp-protocol/tools"
)

// LoginReq 登录
// 说明:
//   - duration传入0, 默认心跳间隔为3s
//   - duration传入0，默认延时器为3s
var LoginReq = func(username string, userId int, duration int) ([]byte, error) {
	tempByte, err := TLogin(username, userId, duration)
	if err != nil {
		return nil, errors.New("Failed to insert into the TLogin template: " + err.Error())
	}
	res := request.Struct2BytesCode(tempByte)
	return res, nil
}

// TLogin Mortal 2026/5/18 16:04 初始化 login 所需字节，返回结构体
var TLogin = func(username string, userId int, duration int) (types.UdpRequestBytesCode, error) {
	res, err := request.TempBase(userId, 0x01)
	if err != nil {
		return nil, err
	}
	bUsername, _ := request.EncodeString(username, "UTF-16BE")
	optCode := tools.Tern(duration > 0, 0x05, 0x01)
	if optCode == 0x01 {
		res = append(res, types.UdpRequestByteCodeItem{
			Name: "Payload",
			Value: types.UdpRequestBytesCode{{
				Name:  "OptCode", // 操作码
				Value: optCode,
				Size:  1,
			}, {
				Name: "OptData",
				Value: types.UdpRequestBytesCode{{
					Name:  "Size", // 用户名的字的个数，最大40个字
					Value: len(username),
					Size:  1,
				}, {
					Name:  "UserName", // UTF-16BE编码。最大长度40个字，用户名不允许重复，如果用户名相同，电台返回用户已登录，由客户端重新命名自己的用户名
					Value: bUsername,
					Size:  len(bUsername),
				}},
				Size: 0,
			}},
			Size: 0,
		})
	} else {
		res = append(res, types.UdpRequestByteCodeItem{
			Name: "Payload",
			Value: types.UdpRequestBytesCode{{
				Name:  "OptCode",
				Value: optCode,
				Size:  1,
			}, {
				Name: "OptData",
				Value: types.UdpRequestBytesCode{{
					Name:  "SuperviseInterval", // 心跳检测间隔。单位：毫秒。最小值：100
					Value: duration,
					Size:  4,
				}, {
					Name:  "SuperviseCnt", // 心跳检测次数
					Value: 3,              // 数字必须 > 2; 0、1、2为无效值
					Size:  2,
				}, {
					Name:  "Size",
					Value: len(username),
					Size:  1,
				}, {
					Name:  "UserName",
					Value: bUsername,
					Size:  len(bUsername),
				}},
				Size: 0,
			}},
			Size: 0,
		})
	}
	return res, nil
}
