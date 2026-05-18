package core

import "github.com/renjietan/hytera-udp-protocol/options"

var ResultPing = func(username string, userId int, alive bool) []byte {
	loginByte, err := TempPing(11)
	if err != nil {
		return nil
	}
	return Struct2Bytes(loginByte)
}

func TempPing(userId int) (options.UdpRequest, error) {
	res, err := TempBase(userId)
	if err != nil {
		return nil, err
	}
	res = append(res, options.Item{
		Name: "Payload",
		Value: options.UdpRequest{{
			Name:  "OptCode",
			Value: 0x03,
			Size:  1,
		}, {
			Name: "OptData",
			Value: options.UdpRequest{{
				Name:  "Status",
				Value: 0x00,
				Size:  1,
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}
