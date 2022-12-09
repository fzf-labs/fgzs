package cache

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var Keys = map[string]*Key{}

// Key 缓存key前缀管理
type Key struct {
	PrefixName     string
	Remark         string
	ExpirationTime time.Duration
}

func NewCacheKey(prefixName string, expirationTime time.Duration, remark string) *Key {
	if _, ok := Keys[prefixName]; ok {
		panic(fmt.Sprintf("cache key %s is exsit, please change one", prefixName))
	}
	key := &Key{PrefixName: prefixName, ExpirationTime: expirationTime, Remark: remark}
	Keys[prefixName] = key
	return key
}

// BuildCacheKey 构建一个带有前缀的缓存key 使用 ":" 分隔
func (p *Key) BuildCacheKey(keys ...string) string {
	if len(keys) == 0 {
		return p.PrefixName
	} else {
		return strings.Join(append([]string{p.PrefixName}, keys...), ":")
	}
}

// TTL 获取缓存key的过期时间
func (p *Key) TTL() time.Duration {
	return p.ExpirationTime
}

func (p *Key) TTLSecond() int {
	return int(p.ExpirationTime / time.Second)
}

func BuildCacheKey(keyPrefix, key string) (cacheKey string, err error) {
	if key == "" {
		return "", errors.New("[cache] key should not be empty")
	}
	cacheKey, err = strings.Join([]string{keyPrefix, key}, ":"), nil
	return
}
