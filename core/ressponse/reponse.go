package ressponse

import (
	"fmt"
	"reflect"

	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types/enums"
)

func ByteCode2Stuct(b []byte, v reflect.Value) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		currentPath := fieldType.Name

		if field.Type() == reflect.TypeOf(reponse.UdpBaseResponse{}) {
			valueField := field.FieldByName("Value")
			sizeField := field.FieldByName("Size")
			if valueField.IsValid() && valueField.CanSet() {
				s := 0
				if sizeField.Int() == 0 {
					s = len(b)
					sizeField.Set(reflect.ValueOf(s))
				} else {
					s = int(sizeField.Int())
				}
				forwardByte := tools.NewSafeBytes(b).Slice(s, enums.Forward)
				backendByte := tools.NewSafeBytes(b).Slice(s, enums.Backward)
				valueField.Set(reflect.ValueOf(forwardByte))
				b = backendByte.Data
				fmt.Printf("%s.Value: %#v\n", currentPath, b)
			}
		} else if field.Kind() == reflect.Struct {
			// 是其他结构体，继续递归
			ByteCode2Stuct(b, field)
		}
	}
}
