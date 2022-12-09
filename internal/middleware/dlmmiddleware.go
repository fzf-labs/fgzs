package middleware

import (
	"bytes"
	"fgzs/internal/define/cachekey"
	"fgzs/internal/errorx"
	"fgzs/internal/meta"
	"fgzs/internal/response"
	"fgzs/pkg/crypt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"io"
	"net/http"
)

type DLMiddleware struct {
	redis *redis.Redis
}

func NewDLMiddleware(redisClient *redis.Redis) *DLMiddleware {
	return &DLMiddleware{redis: redisClient}
}

func (m *DLMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := meta.GetUid(r.Context())
		if uid != "" && r.Method != http.MethodGet {
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				response.Err(r, w, errorx.InternalServerError.WithDetail(err))
				return
			}
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			md5 := crypt.Md5String(string(bodyBytes))
			key := cachekey.DL.BuildCacheKey(uid, r.Method, r.RequestURI, md5)
			redisLock := redis.NewRedisLock(m.redis, key)
			redisLock.SetExpire(cachekey.DL.TTLSecond())
			ok, err := redisLock.Acquire()
			if err != nil {
				logx.Error("DL lock fail")
				response.Err(r, w, errorx.InternalServerError.WithDetail(err))
				return
			}
			if !ok {
				logx.Error("DL lock too fast")
				response.Err(r, w, errorx.RequestFrequencyIsTooFast)
				return
			}
			defer func(redisLock *redis.RedisLock) {
				_, _ = redisLock.Release()
			}(redisLock)
		}
		next(w, r)
	}
}
