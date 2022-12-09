package intutil

import (
	"math/rand"
	"strconv"
	"time"
)

// Int64ToString 将 int64 转换为 string
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// Uint64ToString 将 uint64 转换为 string
func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

// Abs 取绝对值
func Abs(a int) int {
	return (a ^ a>>31) - a>>31
}

// RandInt alias of RandomInt()
func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func RandInt64(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}

// RandIntWithSeed alias of RandomIntWithSeed()
func RandIntWithSeed(min, max int, seed int64) int {
	return RandomIntWithSeed(min, max, seed)
}

// RandomIntWithSeed return a random int at the [min, max)
//
// Usage:
//
//	seed := time.Now().UnixNano()
//	RandomIntWithSeed(1000, 9999, seed)
func RandomIntWithSeed(min, max int, seed int64) int {
	rand.Seed(seed)
	return min + rand.Intn(max-min)
}
