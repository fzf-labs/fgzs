package logic

import (
	"context"
	"fgzs/internal/app/user/rpc/internal/svc"
	"fgzs/internal/app/user/rpc/userpb"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserAppleLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAppleLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAppleLoginLogic {
	return &UserAppleLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录-苹果登录
func (l *UserAppleLoginLogic) UserAppleLogin(in *userpb.UserAppleLoginReq) (*userpb.UserAppleLoginResp, error) {
	return nil, nil
}
