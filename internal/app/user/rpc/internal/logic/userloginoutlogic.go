package logic

import (
	"context"
	"fgzs/internal/app/user/rpc/internal/svc"
	"fgzs/internal/app/user/rpc/userpb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginOutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginOutLogic {
	return &UserLoginOutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登出
func (l *UserLoginOutLogic) UserLoginOut(in *userpb.UserLoginOutReq) (*userpb.UserLoginOutResp, error) {
	// todo: add your logic here and delete this line

	return &userpb.UserLoginOutResp{}, nil
}
