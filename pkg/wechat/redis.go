package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"reflect"
	"strconv"
	"time"
)

type RedisCache struct {
	RedisClient *redis.Redis
}

func NewRedisCache(redisClient *redis.Redis) *RedisCache {
	return &RedisCache{RedisClient: redisClient}
}

func (r *RedisCache) Get(key string) interface{} {
	result, err := r.RedisClient.GetCtx(context.Background(), key)
	if err != nil {
		return nil
	}
	return result
}

func (r *RedisCache) Set(key string, val interface{}, timeout time.Duration) error {
	return r.RedisClient.SetexCtx(context.Background(), key, String(val), int(timeout.Seconds()))
}

func (r *RedisCache) IsExist(key string) bool {
	result, err := r.RedisClient.ExistsCtx(context.Background(), key)
	if err != nil {
		return false
	}
	return result
}

func (r *RedisCache) Delete(key string) error {
	_, err := r.RedisClient.DelCtx(context.Background(), key)
	if err != nil {
		return err
	}
	return nil
}

func String(any interface{}) string {
	if any == nil {
		return ""
	}
	switch value := any.(type) {
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case string:
		return value
	case []byte:
		return string(value)
	case time.Time:
		if value.IsZero() {
			return ""
		}
		return value.String()
	case *time.Time:
		if value == nil {
			return ""
		}
		return value.String()
	default:
		// Empty checks.
		if value == nil {
			return ""
		}
		// Reflect checks.
		var (
			rv   = reflect.ValueOf(value)
			kind = rv.Kind()
		)
		switch kind {
		case reflect.Chan,
			reflect.Map,
			reflect.Slice,
			reflect.Func,
			reflect.Ptr,
			reflect.Interface,
			reflect.UnsafePointer:
			if rv.IsNil() {
				return ""
			}
		case reflect.String:
			return rv.String()
		}
		if kind == reflect.Ptr {
			return String(rv.Elem().Interface())
		}
		// Finally, we use json.Marshal to convert.
		if jsonContent, err := json.Marshal(value); err != nil {
			return fmt.Sprint(value)
		} else {
			return string(jsonContent)
		}
	}
}
