package logic

import (
	"context"
	"fgzs/internal/app/common/rpc/common"
	"fgzs/internal/app/common/rpc/commonpb"
	"fgzs/internal/app/common/rpc/internal/svc"

	"github.com/lithammer/fuzzysearch/fuzzy"

	"github.com/zeromicro/go-zero/core/logx"
)

type SensitiveWordSearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSensitiveWordSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SensitiveWordSearchLogic {
	return &SensitiveWordSearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 敏感词查询
func (l *SensitiveWordSearchLogic) SensitiveWordSearch(in *commonpb.SensitiveWordSearchReq) (*commonpb.SensitiveWordSearchResp, error) {
	sensitiveWordCacheResp, err := NewSensitiveWordCacheLogic(l.ctx, l.svcCtx).SensitiveWordCache(&common.SensitiveWordCacheReq{})
	if err != nil {
		return nil, err
	}
	find := fuzzy.Find(in.Search, sensitiveWordCacheResp.Words)
	return &common.SensitiveWordSearchResp{
		List: find,
	}, nil
}
