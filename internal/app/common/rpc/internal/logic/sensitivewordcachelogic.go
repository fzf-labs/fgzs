package logic

import (
	"context"
	"fgzs/internal/app/common/rpc/commonpb"
	"fgzs/internal/app/common/rpc/internal/svc"
	"fgzs/internal/define/cachekey"
	"fgzs/internal/errorx"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/syncx"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type SensitiveWordCacheLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSensitiveWordCacheLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SensitiveWordCacheLogic {
	return &SensitiveWordCacheLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var SensitiveWordCacheSf = syncx.NewSingleFlight()
var cache *collection.Cache

// 敏感词缓存
func (l *SensitiveWordCacheLogic) SensitiveWordCache(in *commonpb.SensitiveWordCacheReq) (*commonpb.SensitiveWordCacheResp, error) {
	cacheKey := cachekey.SensitiveWord.BuildCacheKey()
	if cache == nil {
		newCache, err := collection.NewCache(time.Minute, collection.WithName(cacheKey))
		if err != nil {
			return nil, err
		}
		cache = newCache
	}
	ret, err := cache.Take(cacheKey, func() (interface{}, error) {
		result, err := l.svcCtx.Redis.SmembersCtx(l.ctx, cacheKey)
		if err != nil && err != redis.Nil {
			return nil, errorx.DataRedisErr.WithDetail(err)
		}
		if len(result) == 0 {
			sfDo, err := SensitiveWordCacheSf.Do(cacheKey, func() (interface{}, error) {
				sensitiveWords, err := l.svcCtx.SensitiveWordModel.FindAllWord(l.ctx)
				if err != nil {
					return nil, err
				}
				wordStrings := make([]string, 0, len(sensitiveWords))
				words := make([]interface{}, 0, len(sensitiveWords))
				if len(sensitiveWords) > 0 {
					for _, v := range sensitiveWords {
						wordStrings = append(wordStrings, v)
						words = append(words, v)
					}
				}
				_, err = l.svcCtx.Redis.SaddCtx(l.ctx, cacheKey, words...)
				if err != nil {
					return nil, errorx.DataRedisErr.WithDetail(err)
				}
				err = l.svcCtx.Redis.ExpireCtx(l.ctx, cacheKey, cachekey.SensitiveWord.TTLSecond())
				if err != nil {
					return nil, errorx.DataRedisErr.WithDetail(err)
				}
				return wordStrings, nil
			})
			if err != nil {
				return nil, err
			}
			result = sfDo.([]string)
		}
		return result, nil
	})
	if err != nil {
		return nil, err
	}
	result := ret.([]string)
	return &commonpb.SensitiveWordCacheResp{Words: result}, nil
}
