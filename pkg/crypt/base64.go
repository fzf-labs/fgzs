package crypt

import "encoding/base64"

// Base64StdEncode 使用 base64 编码对字符串进行编码
func Base64StdEncode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64StdDecode 解码 base64 编码的字符串
func Base64StdDecode(s string) string {
	b, _ := base64.StdEncoding.DecodeString(s)
	return string(b)
}
