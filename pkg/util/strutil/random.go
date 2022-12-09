package strutil

import (
	"math/rand"
	"time"
)

// some consts string chars
const (
	Numbers       = "0123456789"
	AlphaNumLower = "abcdefghijklmnopqrstuvwxyz"
	AlphaNumUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// Random 随机字符串
func Random(n int) string {
	cs := make([]byte, n)
	str := Numbers + AlphaNumLower + AlphaNumUpper
	sl := len(str)
	for i := 0; i < n; i++ {
		// 1607400451937462000
		rand.Seed(time.Now().UnixNano())
		idx := rand.Intn(sl) // 0 - 25
		cs[i] = str[idx]
	}
	return string(cs)
}

// RandomChars 随机字符串
func RandomChars(n int, char ...string) string {
	cs := make([]byte, n)
	str := ""
	if len(char) > 0 {
		for _, s := range char {
			str += s
		}
	} else {
		str = Numbers + AlphaNumLower + AlphaNumUpper
	}
	sl := len(str)
	for i := 0; i < n; i++ {
		// 1607400451937462000
		rand.Seed(time.Now().UnixNano())
		idx := rand.Intn(sl) // 0 - 25
		cs[i] = str[idx]
	}
	return string(cs)
}

// RandomNumber 随机生成指定长度的数字
func RandomNumber(n int) string {
	cs := make([]byte, n)
	str := Numbers
	sl := len(str)
	for i := 0; i < n; i++ {
		// 1607400451937462000
		rand.Seed(time.Now().UnixNano())
		idx := rand.Intn(sl) // 0 - 25
		cs[i] = str[idx]
	}
	return string(cs)
}
