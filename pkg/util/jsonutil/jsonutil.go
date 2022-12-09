package jsonutil

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"regexp"
	"strings"
	"text/scanner"
)

// WriteFile 将数据写入 JSON 文件
func WriteFile(filePath string, data interface{}) error {
	jsonBytes, err := Encode(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, jsonBytes, 0664)
}

// ReadFile 读取 JSON 文件数据
func ReadFile(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()
	return json.NewDecoder(file).Decode(v)
}

// Pretty json 友好打印
func Pretty(v interface{}) (string, error) {
	out, err := json.MarshalIndent(v, "", "    ")
	return string(out), err
}

// Encode 编码
func Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// EncodePretty 编码为友好的 JSON
func EncodePretty(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "    ")
}

// EncodeToWriter 编码到 io.Writer
func EncodeToWriter(v interface{}, w io.Writer) error {
	return json.NewEncoder(w).Encode(v)
}

// EncodeUnescapeHTML 编码,去除html转义
func EncodeUnescapeHTML(v interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decode 解码
func Decode(bts []byte, ptr interface{}) error {
	return json.Unmarshal(bts, ptr)
}

// DecodeString 解码字符串
func DecodeString(str string, ptr interface{}) error {
	return json.Unmarshal([]byte(str), ptr)
}

// DecodeReader 解码 io.Reader
func DecodeReader(r io.Reader, ptr interface{}) error {
	return json.NewDecoder(r).Decode(ptr)
}

// Mapping src data(map,struct) to dst struct use json tags.
//
// On src, dst both is struct, equivalent to merging two structures (src should be a subset of dsc)
func Mapping(src, dst interface{}) error {
	bts, err := Encode(src)
	if err != nil {
		return err
	}
	return Decode(bts, dst)
}

// `(?s:` enable match multi line
var jsonMLComments = regexp.MustCompile(`(?s:/\*.*?\*/\s*)`)

// StripComments 去除 JSON 字符串的注释
func StripComments(src string) string {
	// multi line comments
	if strings.Contains(src, "/*") {
		src = jsonMLComments.ReplaceAllString(src, "")
	}

	// single line comments
	if !strings.Contains(src, "//") {
		return strings.TrimSpace(src)
	}

	// strip inline comments
	var s scanner.Scanner

	s.Init(strings.NewReader(src))
	s.Filename = "comments"
	s.Mode ^= scanner.SkipComments // don't skip comments

	buf := new(bytes.Buffer)
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		txt := s.TokenText()
		if !strings.HasPrefix(txt, "//") && !strings.HasPrefix(txt, "/*") {
			buf.WriteString(txt)
			// } else {
			// fmt.Printf("%s: %s\n", s.Position, txt)
		}
	}

	return buf.String()
}
