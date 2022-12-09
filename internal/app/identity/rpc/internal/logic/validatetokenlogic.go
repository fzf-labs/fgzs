package logic

import (
	"context"
	"fgzs/internal/app/identity/rpc/identity"
	"fgzs/internal/app/identity/rpc/identitypb"
	"fgzs/internal/app/identity/rpc/internal/svc"
	"fgzs/internal/errorx"
	"fgzs/pkg/conv"
	"fgzs/pkg/jwt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type ValidateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateTokenLogic {
	return &ValidateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证token
func (l *ValidateTokenLogic) ValidateToken(in *identitypb.ValidateTokenReq) (*identitypb.ValidateTokenResp, error) {
	//配置获取
	jwtConfig, ok := l.svcCtx.Config.Jwt[in.Target]
	if !ok {
		return nil, errorx.TokenWrongTypeOfBusiness
	}
	//Token 解析
	j := jwt.NewJwt(l.svcCtx.Redis, jwtConfig)
	claims, err := j.ParseToken(in.Token)
	if err != nil {
		return nil, errorx.TokenInvalidErr.WithDetail(err)
	}
	//校验黑名单
	blackTokenCheck, newToken := j.JwtBlackTokenCheck(claims)
	if blackTokenCheck {
		newClaims, err := j.ParseToken(newToken.Token)
		if err != nil {
			return nil, errorx.TokenInvalidErr.WithDetail(err)
		}
		return &identity.ValidateTokenResp{
			Uid:      conv.String(newClaims[jwt.JwtUID]),
			Payloads: j.GetPayloads(newClaims),
			Token: &identity.Token{
				Token:     newToken.Token,
				ExpiredAt: newToken.ExpiredAt,
				RefreshAt: newToken.RefreshAt,
			},
		}, nil
	}
	//校验是否单一登录
	err = j.JwtTokenCheck(claims)
	if err != nil {
		return nil, errorx.TokenVerificationFailed
	}
	//校验是否可以刷新
	if conv.Int64(claims[jwt.JwtRefresh]) > time.Now().Unix() {
		newToken, newClaims, err := j.RefreshToken(claims)
		if err != nil {
			return nil, err
		}
		return &identity.ValidateTokenResp{
			Uid:      conv.String(newClaims[jwt.JwtUID]),
			Payloads: j.GetPayloads(newClaims),
			Token: &identity.Token{
				Token:     newToken.Token,
				ExpiredAt: newToken.ExpiredAt,
				RefreshAt: newToken.RefreshAt,
			},
		}, nil
	}
	return &identity.ValidateTokenResp{
		Uid:      conv.String(claims[jwt.JwtUID]),
		Payloads: j.GetPayloads(claims),
		Token:    nil,
	}, nil
}
