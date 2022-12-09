package logic

import (
	"context"
	"fgzs/internal/app/user/rpc/internal/svc"
	"fgzs/internal/app/user/rpc/userpb"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserUsernameLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUsernameLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUsernameLoginLogic {
	return &UserUsernameLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录-用户名密码
func (l *UserUsernameLoginLogic) UserUsernameLogin(in *userpb.UserUsernameLoginReq) (*userpb.UserUsernameLoginResp, error) {
	return nil, nil
}
