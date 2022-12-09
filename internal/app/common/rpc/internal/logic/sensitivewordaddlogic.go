package logic

import (
	"context"
	"fgzs/internal/app/common/model"
	"fgzs/internal/app/common/rpc/common"
	"fgzs/internal/app/common/rpc/commonpb"
	"fgzs/internal/app/common/rpc/internal/svc"
	"fgzs/internal/define/cachekey"
	"fgzs/internal/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type SensitiveWordAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSensitiveWordAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SensitiveWordAddLogic {
	return &SensitiveWordAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 敏感词添加
func (l *SensitiveWordAddLogic) SensitiveWordAdd(in *commonpb.SensitiveWordAddReq) (*commonpb.SensitiveWordAddResp, error) {
	sensitiveWordInfo, err := l.svcCtx.SensitiveWordModel.FindOneByWord(l.ctx, in.Word)
	if err != nil {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	if sensitiveWordInfo != nil {
		return nil, errorx.DataDuplicateRecords
	}
	_, err = l.svcCtx.SensitiveWordModel.Insert(l.ctx, &model.SensitiveWord{
		Word: in.Word,
	})
	if err != nil {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	cacheKey := cachekey.SensitiveWord.BuildCacheKey()
	exists, err := l.svcCtx.Redis.ExistsCtx(l.ctx, cacheKey)
	if err != nil {
		return nil, errorx.DataRedisErr.WithDetail(err)
	}
	if exists {
		_, err := l.svcCtx.Redis.SaddCtx(l.ctx, cacheKey, in.Word)
		if err != nil {
			return nil, errorx.DataRedisErr.WithDetail(err)
		}
	}
	return &common.SensitiveWordAddResp{}, nil
}
