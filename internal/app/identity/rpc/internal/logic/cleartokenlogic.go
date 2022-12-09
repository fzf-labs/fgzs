package logic

import (
	"context"
	"fgzs/internal/app/identity/rpc/identitypb"
	"fgzs/internal/app/identity/rpc/internal/svc"
	"fgzs/internal/errorx"
	"fgzs/pkg/jwt"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClearTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClearTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearTokenLogic {
	return &ClearTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 清除token
func (l *ClearTokenLogic) ClearToken(in *identitypb.ClearTokenReq) (*identitypb.ClearTokenResp, error) {
	//配置获取
	jwtConfig, ok := l.svcCtx.Config.Jwt[in.Target]
	if !ok {
		return nil, errorx.TokenWrongTypeOfBusiness
	}
	err := jwt.NewJwt(l.svcCtx.Redis, jwtConfig).JwtTokenClear(in.Uid)
	if err != nil {
		return nil, err
	}
	return &identitypb.ClearTokenResp{}, nil
}
