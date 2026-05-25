package response

import (
	"fmt"
	"reflect"

	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types/enums"
	types "github.com/renjietan/hytera-udp-protocol/types/reponse"
)

// ByteCode2Stuct
// 使用方式
//
//	var b = []byte{
//		0xee,      // SrcID
//		0xee,      // SrcID
//		0x0, 0x17, // Length
//		0x0, 0x0, // CRC
//		0x0, // Version
//		0xb, // UserID
//		0x1, // SAP
//		0x5, // OptCode
//		0x6, // Size
//		0x7, // OptCode
//		0x8, // Size
//		0x9} // UserId
//
// loginAck := core_response_auth.LoginRes()
// response.ByteCode2Stuct(b, &loginAck, "")
// fmt.Println("loginAck=========", loginAck)
func ByteCode2Stuct(b []byte, v interface{}, path string) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		currentPath := path + fieldType.Name
		if fieldType.Type == reflect.TypeOf(types.UdpResponseByteCodeItem{}) {
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
				valueField.Set(reflect.ValueOf(forwardByte.Data))
				b = backendByte.Data
			}
			fmt.Println(currentPath)
			continue
		}

		if fieldType.Type.Kind() == reflect.Interface {
			if field.IsNil() {
				continue
			}
			elem := field.Elem()
			if elem.Kind() == reflect.Struct {
				newPtr := reflect.New(elem.Type())
				newPtr.Elem().Set(elem)
				ByteCode2Stuct(b, newPtr.Interface(), currentPath+".")
				field.Set(newPtr.Elem())
			} else if elem.Kind() == reflect.Ptr {
				ByteCode2Stuct(b, elem.Interface(), currentPath+".")
			}
			continue
		}
		if field.Kind() == reflect.Struct {
			if field.CanAddr() {
				ByteCode2Stuct(b, field.Addr().Interface(), currentPath+".")
			}
		}
	}
}

//func ByteCode2Stuct(b []byte, v reflect.Value) {
//	if v.Kind() == reflect.Ptr {
//		v = v.Elem()
//	}
//	if v.Kind() != reflect.Struct {
//		return
//	}
//	t := v.Type()
//	for i := 0; i < v.NumField(); i++ {
//		field := v.Field(i)
//		fieldType := t.Field(i)
//		currentPath := fieldType.Name
//
//		if field.Type() == reflect.TypeOf(types.UdpResponseByteCodeItem{}) {
//			valueField := field.FieldByName("Value")
//			sizeField := field.FieldByName("Size")
//			if valueField.IsValid() && valueField.CanSet() {
//				s := 0
//				if sizeField.Int() == 0 {
//					s = len(b)
//					sizeField.Set(reflect.ValueOf(s))
//				} else {
//					s = int(sizeField.Int())
//				}
//				forwardByte := tools.NewSafeBytes(b).Slice(s, enums.Forward)
//				backendByte := tools.NewSafeBytes(b).Slice(s, enums.Backward)
//				valueField.Set(reflect.ValueOf(forwardByte))
//				b = backendByte.Data
//				fmt.Printf("%s.Value: %#v\n", currentPath, b)
//			}
//		} else if field.Elem().Kind() == reflect.Struct {
//			ByteCode2Stuct(b, field)
//		}
//	}
//}
