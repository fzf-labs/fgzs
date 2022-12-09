package strutil

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

// StrToLower 转换成小写字母
func StrToLower(str string) string {
	runeArr := []rune(str)
	for i := range runeArr {
		if runeArr[i] >= 65 && runeArr[i] <= 90 {
			runeArr[i] += 32
		}
	}
	return string(runeArr)
}

// ConcatString 连接字符串
// NOTE: 性能比fmt.Sprintf和+号要好
func ConcatString(s ...string) string {
	if len(s) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	for _, i := range s {
		buffer.WriteString(i)
	}
	return buffer.String()
}

// StringToUint64 字符串转uint64
func StringToUint64(str string) (uint64, error) {
	if str == "" {
		return 0, nil
	}
	valInt, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return uint64(valInt), nil
}

// StringToInt64 字符串转int64
func StringToInt64(str string) (int64, error) {
	if str == "" {
		return 0, nil
	}
	valInt, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return int64(valInt), nil
}

// StringToInt 字符串转int
func StringToInt(str string) (int, error) {
	if str == "" {
		return 0, nil
	}
	valInt, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return valInt, nil
}

// --------- 字节切片和字符串转换 ----------
// 性能很高, 原因在于底层无新的内存申请与拷贝

// Bytes2String 字节切片转字符串
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// String2Bytes 字符串转字节切片
func String2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// SubStr 截取字符串，并返回实际截取的长度和子串
func SubStr(str string, start, end int64) (int64, string, error) {
	reader := strings.NewReader(str)

	// Calling NewSectionReader method with its parameters
	r := io.NewSectionReader(reader, start, end)

	// Calling Copy method with its parameters
	var buf bytes.Buffer
	n, err := io.Copy(&buf, r)
	return n, buf.String(), err
}

// SubstrTarget 在字符串中查找指定子串，并返回left或right部分
func SubstrTarget(str string, target string, turn string, hasPos bool) (string, error) {
	pos := strings.Index(str, target)

	if pos == -1 {
		return "", nil
	}

	if turn == "left" {
		if hasPos {
			pos = pos + 1
		}
		return str[:pos], nil
	} else if turn == "right" {
		if !hasPos {
			pos = pos + 1
		}
		return str[pos:], nil
	} else {
		return "", errors.New("params 3 error")
	}
}

// GetStringUtf8Len 获得字符串按照uft8编码的长度
func GetStringUtf8Len(str string) int {
	return utf8.RuneCountInString(str)
}

// Utf8Index 按照uft8编码匹配子串，返回开头的索引
func Utf8Index(str, substr string) int {
	index := strings.Index(str, substr)
	if index < 0 {
		return -1
	}
	return utf8.RuneCountInString(str[:index])
}

// JoinStringAndOther 连接字符串和其他类型
// fmt.Println(JoinStringAndOther("why", 123))
func JoinStringAndOther(val ...interface{}) string {
	return fmt.Sprint(val...)
}

// UcFirst 首字母大写
func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// LcFirst 首字母小写
func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// FormatPrivateKey 格式化 普通应用秘钥
func FormatPrivateKey(privateKey string) (pKey string) {
	var buffer strings.Builder
	buffer.WriteString("-----BEGIN RSA PRIVATE KEY-----\n")
	rawLen := 64
	keyLen := len(privateKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(privateKey[start:])
		} else {
			buffer.WriteString(privateKey[start:end])
		}
		buffer.WriteByte('\n')
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END RSA PRIVATE KEY-----\n")
	pKey = buffer.String()
	return
}

// FormatPublicKey 格式化 普通支付宝公钥
func FormatPublicKey(publicKey string) (pKey string) {
	var buffer strings.Builder
	buffer.WriteString("-----BEGIN PUBLIC KEY-----\n")
	rawLen := 64
	keyLen := len(publicKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(publicKey[start:])
		} else {
			buffer.WriteString(publicKey[start:end])
		}
		buffer.WriteByte('\n')
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END PUBLIC KEY-----\n")
	pKey = buffer.String()
	return
}

// CamelCase covert string to camelCase string.
func CamelCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	result := ""
	blankSpace := " "
	regex, _ := regexp.Compile("[-_&]+")
	ss := regex.ReplaceAllString(s, blankSpace)
	for i, v := range strings.Split(ss, blankSpace) {
		vv := []rune(v)
		if i == 0 {
			if vv[i] >= 65 && vv[i] <= 96 {
				vv[0] += 32
			}
			result += string(vv)
		} else {
			result += Capitalize(v)
		}
	}

	return result
}

// Capitalize converts the first character of a string to upper case and the remaining to lower case.
func Capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}

	out := make([]rune, len(s))
	for i, v := range s {
		if i == 0 {
			out[i] = unicode.ToUpper(v)
		} else {
			out[i] = unicode.ToLower(v)
		}
	}

	return string(out)
}

// KebabCase covert string to kebab-case
func KebabCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	regex := regexp.MustCompile(`[\W|_]+`)
	blankSpace := " "
	match := regex.ReplaceAllString(s, blankSpace)
	rs := strings.Split(match, blankSpace)

	var result []string
	for _, v := range rs {
		splitWords := splitWordsToLower(v)
		if len(splitWords) > 0 {
			result = append(result, splitWords...)
		}
	}

	return strings.Join(result, "-")
}

// SnakeCase covert string to snake_case
func SnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	regex := regexp.MustCompile(`[\W|_]+`)
	blankSpace := " "
	match := regex.ReplaceAllString(s, blankSpace)
	rs := strings.Split(match, blankSpace)

	var result []string
	for _, v := range rs {
		splitWords := splitWordsToLower(v)
		if len(splitWords) > 0 {
			result = append(result, splitWords...)
		}
	}

	return strings.Join(result, "_")
}

// splitWordsToLower split a string into worlds by uppercase char
func splitWordsToLower(s string) []string {
	var result []string

	upperIndexes := upperIndex(s)
	l := len(upperIndexes)
	if upperIndexes == nil || l == 0 {
		if s != "" {
			result = append(result, s)
		}
		return result
	}
	for i := 0; i < l; i++ {
		if i < l-1 {
			result = append(result, strings.ToLower(s[upperIndexes[i]:upperIndexes[i+1]]))
		} else {
			result = append(result, strings.ToLower(s[upperIndexes[i]:]))
		}
	}
	return result
}

// upperIndex get a int slice which elements are all the uppercase char index of a string
func upperIndex(s string) []int {
	var result []int
	for i := 0; i < len(s); i++ {
		if 64 < s[i] && s[i] < 91 {
			result = append(result, i)
		}
	}
	if len(s) > 0 && result != nil && result[0] != 0 {
		result = append([]int{0}, result...)
	}

	return result
}

// Reverse return string whose char order is reversed to the given string
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
