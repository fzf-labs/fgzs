package logic

import (
	"context"
	"fgzs/internal/app/common/rpc/common"
	"fgzs/internal/app/common/rpc/commonpb"
	"fgzs/internal/app/common/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SmsSendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSmsSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SmsSendLogic {
	return &SmsSendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 短信发送
func (l *SmsSendLogic) SmsSend(in *commonpb.SmsSendReq) (*commonpb.SmsSendResp, error) {
	return &common.SmsSendResp{}, nil
}
