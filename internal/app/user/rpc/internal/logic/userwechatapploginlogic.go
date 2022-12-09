package logic

import (
	"context"
	"fgzs/internal/app/user/rpc/internal/svc"
	"fgzs/internal/app/user/rpc/userpb"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserWechatAppLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserWechatAppLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserWechatAppLoginLogic {
	return &UserWechatAppLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录-微信登录
func (l *UserWechatAppLoginLogic) UserWechatAppLogin(in *userpb.UserWechatAppLoginReq) (*userpb.UserWechatAppLoginResp, error) {
	return nil, nil
}
