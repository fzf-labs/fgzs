package logic

import (
	"context"
	"fgzs/internal/app/user/rpc/internal/svc"
	"fgzs/internal/app/user/rpc/userpb"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserSmsLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserSmsLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSmsLoginLogic {
	return &UserSmsLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录-短信登录
func (l *UserSmsLoginLogic) UserSmsLogin(in *userpb.UserSmsLoginReq) (*userpb.UserSmsLoginResp, error) {
	return nil, nil
}
