package logic

import (
	"context"
	"fgzs/internal/app/user/rpc/internal/svc"
	"fgzs/internal/app/user/rpc/userpb"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserSmsSendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserSmsSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSmsSendLogic {
	return &UserSmsSendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录-短信发送
func (l *UserSmsSendLogic) UserSmsSend(in *userpb.UserSmsSendReq) (*userpb.UserSmsSendResp, error) {
	return nil, nil
}
