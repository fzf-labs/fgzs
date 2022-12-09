package logic

import (
	"context"
	"fgzs/internal/app/common/rpc/common"
	"fgzs/internal/app/common/rpc/commonpb"
	"fgzs/internal/app/common/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SensitiveWordDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSensitiveWordDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SensitiveWordDelLogic {
	return &SensitiveWordDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 敏感词删除
func (l *SensitiveWordDelLogic) SensitiveWordDel(in *commonpb.SensitiveWordDelReq) (*commonpb.SensitiveWordDelResp, error) {
	err := l.svcCtx.SensitiveWordModel.DeleteByWord(l.ctx, in.Word)
	if err != nil {
		return nil, err
	}
	return &common.SensitiveWordDelResp{}, nil
}
