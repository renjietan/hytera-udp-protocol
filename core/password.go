package core

import (
	"errors"

	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types"
)

var Password = func(password string, userId int) ([]byte, error) {
	tempByte, err := TempLogout(password, userId)
	if err != nil {
		return nil, errors.New("Failed to insert into the password template: " + err.Error())
	}
	recordField := tools.GetRecursiveField(tempByte, []types.Item{})
	_, res := tools.Struct2Bytes(tempByte, recordField, []byte{})
	return res, nil
}

var TempPassword = func(password string, userId int) (types.UdpRequest, error) {
	res, err := TempBase(userId, 0x01)
	if err != nil {
		return nil, err
	}
	bPwd, _ := tools.EncodeString(password, "UTF-16BE")
	res = append(res, types.Item{
		Name: "Payload",
		Value: types.UdpRequest{{
			Name:  "OptCode",
			Value: 0x04,
			Size:  1,
		}, {
			Name: "OptData",
			Value: types.UdpRequest{{
				Name:  "Size",
				Value: len(password),
				Size:  1,
			}, {
				Name:  "UserName",
				Value: bPwd,
				Size:  len(bPwd),
			}},
			Size: 0,
		}},
		Size: 0,
	})
	return res, nil
}
