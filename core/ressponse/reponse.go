package response

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"strings"

	types "github.com/renjietan/hytera-udp-protocol/core/ressponse/types"
	"github.com/renjietan/hytera-udp-protocol/tools"
	"github.com/renjietan/hytera-udp-protocol/types/enums"
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

// ByteCode2Stuct 码流转结构体
// params[b] 码流
// params[v] 码流转换后的结构体指针，例如 lTemp := core_response_auth.LoginRes()
// params[path] 当前循环的路径，例如 Payload.OptData.Status
// params[bindPaths] 例如: 递归到 Payload.OptData.UserName 时，若 bindPaths 中存在该路径，则取出对应数据赋给 UserName 的 Size 字段
func ByteCode2Stuct(b []byte, v interface{}, path string, bindPaths map[string]types.UdpResponseBindStruct) {
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
			bindField := field.FieldByName("Bind")
			rawBindField := bindField.Interface()
			originalBind, _ := rawBindField.(types.UdpResponseBindStruct)

			if valueField.IsValid() && valueField.CanSet() {
				s := 0
				if sizeField.Int() == 0 {
					s = len(b)
					sizeField.Set(reflect.ValueOf(s))
				} else {
					s = int(sizeField.Int())
				}
				// 循环到 bind 的字段时，重新赋值
				value := bindPaths[currentPath]
				Path := value.Path
				ValueType := value.ValueType
				FieldName := value.FieldName
				if Path != "" {
					if ValueType == types.UdpResponseBindStructInt {
						s = value.Value.(int)
						if FieldName == types.UdpResponseBindStructSize {
							sizeField.Set(reflect.ValueOf(s))
						}
					}
				}
				forwardByte := tools.NewSafeBytes(b).Slice(s, enums.Forward)
				backendByte := tools.NewSafeBytes(b).Slice(s, enums.Backward)
				valueField.Set(reflect.ValueOf(forwardByte.Data))
				// 存在 bind  字段时，存下来
				if originalBind.Path != "" {
					if originalBind.ValueType == types.UdpResponseBindStructString {
						originalBind.Value = string(forwardByte.Data)
					} else if originalBind.ValueType == types.UdpResponseBindStructInt {
						pathSlice := strings.Split(originalBind.Path, ".")
						ext := pathSlice[len(pathSlice)-1]
						if ext == "UserName" {
							originalBind.Value = int(BytesToUintBE(forwardByte.Data)) * 2
						} else {
							originalBind.Value = int(BytesToUintBE(forwardByte.Data))
						}
					}
					bindPaths[originalBind.Path] = originalBind
				}
				fmt.Println(currentPath, "===================", bindPaths)
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
				ByteCode2Stuct(b, newPtr.Interface(), currentPath+".", bindPaths)
				field.Set(newPtr.Elem())
			} else if elem.Kind() == reflect.Ptr {
				ByteCode2Stuct(b, elem.Interface(), currentPath+".", bindPaths)
			}
			continue
		}
		if field.Kind() == reflect.Struct {
			if field.CanAddr() {
				ByteCode2Stuct(b, field.Addr().Interface(), currentPath+".", bindPaths)
			}
		}
	}
}

func BytesToUintBE(b []byte) uint64 {
	switch len(b) {
	case 1:
		return uint64(b[0])
	case 2:
		return uint64(binary.BigEndian.Uint16(b))
	case 4:
		return uint64(binary.BigEndian.Uint32(b))
	case 8:
		return binary.BigEndian.Uint64(b)
	default:
		return 0
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
