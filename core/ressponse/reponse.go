package response

import (
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
// params[bindPaths.Path] 关联字段的路径，例如 Payload.OptData.UserName
// params[bindPaths.FieldName] 给字段的哪个属性赋值  例如 Size | Value
// params[bindPaths.Value] 将 Value 值 赋值给 Payload.OptData.UserName
// params[bindPaths.ValueType] Value 值类型，如果是List，则需要生成列表， int 或 string  则需要转换
// 例如: 递归时，
// 1、当 bindPaths 存在 [currentPath] 的键名时，则将 [bindPaths.Value] 的值，赋值给当前字段的 [bindPaths FieldName] 属性
// 2、当 参数 v 的 元素中存在 [bind] 字段时，根据 [bind.ValueType] 来执行接下来的操作，根据 [bind.FieldName] 来针对 某个字段进行赋值，一般是 Size 或 Value
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

		value := bindPaths[currentPath]
		Path := value.Path
		ValueType := value.ValueType
		FieldName := value.FieldName
		Value := value.Value
		if fieldType.Type == reflect.TypeOf([]types.AdapterInfoItem{}) {
			m := reflect.MakeSlice(reflect.TypeOf([]types.AdapterInfoItem{}), 0, 0)
			for _, e := range Value.([]types.AdapterInfoItem) {
				m = reflect.Append(m, reflect.ValueOf(e))
			}
			field.Set(m)
		} else if fieldType.Type == reflect.TypeOf(types.UdpResponseByteCodeItem{}) {
			valueField := field.FieldByName("Value")
			sizeField := field.FieldByName("Size")
			bindField := field.FieldByName("Bind")
			rawBindField := bindField.Interface()
			originalBind, _ := rawBindField.(types.UdpResponseBindStruct)

			if valueField.IsValid() && valueField.CanSet() {
				s := 0
				if sizeField.Int() != 0 {
					s = int(sizeField.Int())
				}
				// 循环到 bind 的字段时，重新赋值
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
				b = backendByte.Data
				// 存在 bind  字段时，存下来
				if originalBind.Path != "" {
					if originalBind.ValueType == types.UdpResponseBindStructString {
						originalBind.Value = string(forwardByte.Data)
					} else if originalBind.ValueType == types.UdpResponseBindStructInt {
						ext := GetPathExt(originalBind.Path, ".")
						// UserName 是 UTF16BE
						if ext == "UserName" {
							originalBind.Value = int(tools.BytesToUintBE(forwardByte.Data)) * 2
						} else {
							originalBind.Value = int(tools.BytesToUintBE(forwardByte.Data))
						}
					} else if originalBind.ValueType == types.UdpResponseBindStructRange {
						ext := GetPathExt(originalBind.Path, ".")
						itemLen := int(tools.BytesToUintBE(forwardByte.Data))

						//vv := tools.BytesToUintBE(valueField)
						if ext == "Channels" {
							originalValue, bt := NewAdapterInfoItems(itemLen, b)
							b = bt
							originalBind.Value = originalValue
						}
					}
					bindPaths[originalBind.Path] = originalBind
				}

			}
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

func GetPathExt(path string, splitStr string) string {
	pathSlice := strings.Split(path, splitStr)
	ext := pathSlice[len(pathSlice)-1]
	return ext
}

func NewAdapterInfoItems(itemLen int, b []byte) ([]types.AdapterInfoItem, []byte) {
	var m []types.AdapterInfoItem
	for _ = range itemLen {
		ChannelNo := types.UdpResponseByteCodeItem{
			Value: b[:1],
			Size:  1,
		}
		b = b[1:]
		AdapterType := types.UdpResponseByteCodeItem{
			Value: b[:1],
			Size:  1,
		}
		b = b[1:]
		Ipv4Addr := types.UdpResponseByteCodeItem{
			Value: b[:4],
			Size:  4,
		}
		b = b[4:]
		Ipv4Mask := types.UdpResponseByteCodeItem{
			Value: b[:4],
			Size:  4,
		}
		b = b[4:]
		SValue := tools.BytesToUintBE(b[:1])
		Size := types.UdpResponseByteCodeItem{
			Value: SValue,
			Size:  1,
		}
		b = b[1:]
		AdapterName := types.UdpResponseByteCodeItem{
			Value: b[:SValue*2],
			Size:  int(SValue * 2),
		}
		b = b[SValue*2:]
		m = append(m, types.AdapterInfoItem{
			ChannelNo:   ChannelNo,
			AdapterType: AdapterType,
			Ipv4Addr:    Ipv4Addr,
			Ipv4Mask:    Ipv4Mask,
			Size:        Size,
			AdapterName: AdapterName,
		})
	}
	return m, b
}
