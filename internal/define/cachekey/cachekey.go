package cachekey

import (
	"fgzs/pkg/cache"
	"time"
)

// 缓存key前缀
var (
	UUID = cache.NewCacheKey("uuid", time.Hour, "uuid")
	DL   = cache.NewCacheKey("dl", time.Second*5, "分布式锁")

	Sms       = cache.NewCacheKey("sms", time.Minute*5, "短信验证")
	SmsDayNum = cache.NewCacheKey("sms_day_num", time.Minute*5, "短信发送次数")

	SensitiveWord = cache.NewCacheKey("sensitive_word", time.Hour*24, "敏感词")
)
