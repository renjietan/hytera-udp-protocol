package tools

import (
	"encoding/binary"
	"strings"
	"unicode/utf16"
)

// EncodeString 将字符串转换为指定编码的字节切片
// 支持的编码：
//   - UTF-16BE    (大端序 UTF-16)
//   - UTF-16LE    (小端序 UTF-16)
//   - UTF-32BE    (大端序 UTF-32)
//   - UTF-32LE    (小端序 UTF-32)
func EncodeString(s string, encoding string) ([]byte, error) {
	encoding = strings.ToUpper(encoding)

	switch encoding {
	case "UTF-16BE":
		// 将字符串转为 UTF-16 码点序列（使用标准库 unicode/utf16）
		runes := []rune(s)                     // 转为 rune (Unicode 码点)
		utf16Surrogates := utf16.Encode(runes) // 转为 UTF-16 代理对序列
		buf := make([]byte, 2*len(utf16Surrogates))
		for i, v := range utf16Surrogates {
			binary.BigEndian.PutUint16(buf[2*i:], v)
		}
		return buf, nil

	case "UTF-16LE":
		runes := []rune(s)
		utf16Surrogates := utf16.Encode(runes)
		buf := make([]byte, 2*len(utf16Surrogates))
		for i, v := range utf16Surrogates {
			binary.LittleEndian.PutUint16(buf[2*i:], v)
		}
		return buf, nil

	case "UTF-32BE":
		// UTF-32 直接使用 rune (4字节)
		runes := []rune(s)
		buf := make([]byte, 4*len(runes))
		for i, r := range runes {
			binary.BigEndian.PutUint32(buf[4*i:], uint32(r))
		}
		return buf, nil

	case "UTF-32LE":
		runes := []rune(s)
		buf := make([]byte, 4*len(runes))
		for i, r := range runes {
			binary.LittleEndian.PutUint32(buf[4*i:], uint32(r))
		}
		return buf, nil

	default:
		return nil, nil
	}
}
