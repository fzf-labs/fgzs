package logic

import (
	"context"
	"fgzs/internal/app/common/rpc/common"
	"fgzs/internal/app/common/rpc/commonpb"
	"fgzs/internal/app/common/rpc/internal/svc"
	"sync"

	"github.com/importcjj/sensitive"
	"github.com/zeromicro/go-zero/core/logx"
)

type SensitiveWordCheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSensitiveWordCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SensitiveWordCheckLogic {
	return &SensitiveWordCheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type SensitiveCheck struct {
	len    int
	filter *sensitive.Filter
}

var lock sync.Mutex
var sc = new(SensitiveCheck)

// 敏感词检测
func (l *SensitiveWordCheckLogic) SensitiveWordCheck(in *commonpb.SensitiveWordCheckReq) (*commonpb.SensitiveWordCheckResp, error) {
	sensitiveWordCacheResp, err := NewSensitiveWordCacheLogic(l.ctx, l.svcCtx).SensitiveWordCache(&common.SensitiveWordCacheReq{})
	if err != nil {
		return nil, err
	}
	if sc.len != len(sensitiveWordCacheResp.Words) {
		lock.Lock()
		defer lock.Unlock()
		sc = &SensitiveCheck{
			len:    len(sensitiveWordCacheResp.Words),
			filter: sensitive.New(),
		}
		sc.filter.AddWord(sensitiveWordCacheResp.Words...)
		sc.filter.UpdateNoisePattern(`x`)
	}
	validate, _ := sc.filter.Validate(in.Word)
	if validate {
		return &common.SensitiveWordCheckResp{
			Result:  false,
			Replace: "",
			Filter:  "",
		}, nil
	}
	replace := sc.filter.Replace(in.Word, '*')
	filterStr := sc.filter.Filter(in.Word)
	return &common.SensitiveWordCheckResp{
		Result:  true,
		Replace: replace,
		Filter:  filterStr,
	}, nil
}
