package logic

import (
	"context"
	"fgzs/internal/app/user/rpc/internal/svc"
	"fgzs/internal/app/user/rpc/userpb"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户注册
func (l *UserRegisterLogic) UserRegister(in *userpb.UserRegisterReq) (*userpb.UserRegisterResp, error) {
	return nil, nil
}
