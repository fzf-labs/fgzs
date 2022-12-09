package logic

import (
	"context"
	"fgzs/internal/app/user/rpc/internal/svc"
	"fgzs/internal/app/user/rpc/userpb"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserQQAppLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserQQAppLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserQQAppLoginLogic {
	return &UserQQAppLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录-qq登录
func (l *UserQQAppLoginLogic) UserQQAppLogin(in *userpb.UserQQAppLoginReq) (*userpb.UserQQAppLoginResp, error) {
	return nil, nil
}
