package logic

import (
	"context"
	"fgzs/internal/app/identity/rpc/identity"
	"fgzs/internal/app/identity/rpc/identitypb"
	"fgzs/internal/app/identity/rpc/internal/svc"
	"fgzs/internal/errorx"
	"fgzs/pkg/jwt"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 生成token
func (l *GenerateTokenLogic) GenerateToken(in *identitypb.GenerateTokenReq) (*identitypb.GenerateTokenResp, error) {
	//配置获取
	jwtConfig, ok := l.svcCtx.Config.Jwt[in.Target]
	if !ok {
		return nil, errorx.TokenWrongTypeOfBusiness
	}
	kv := make(map[string]interface{})
	kv["uid"] = in.Uid
	if len(in.Payloads) > 0 {
		for k := range in.Payloads {
			kv[k] = in.Payloads[k]
		}
	}
	jwtClient := jwt.NewJwt(l.svcCtx.Redis, jwtConfig)
	token, claims, err := jwtClient.GenerateToken(kv)
	if err != nil {
		return nil, err
	}
	err = jwtClient.JwtTokenStore(claims)
	if err != nil {
		return nil, errorx.TokenStorageFailed.WithDetail(err)
	}
	return &identity.GenerateTokenResp{Token: &identity.Token{
		Token:     token.Token,
		ExpiredAt: token.ExpiredAt,
		RefreshAt: token.RefreshAt,
	}}, nil
}
