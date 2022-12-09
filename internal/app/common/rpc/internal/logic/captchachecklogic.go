package logic

import (
	"context"
	"fgzs/internal/app/common/rpc/common"
	"fgzs/internal/app/common/rpc/commonpb"
	"fgzs/internal/app/common/rpc/internal/svc"
	"fgzs/internal/errorx"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaCheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCaptchaCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaCheckLogic {
	return &CaptchaCheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证码校验
func (l *CaptchaCheckLogic) CaptchaCheck(in *commonpb.CaptchaCheckReq) (*commonpb.CaptchaCheckResp, error) {
	verify := base64Captcha.DefaultMemStore.Verify(in.CaptchaId, in.Captcha, true)
	if !verify {
		return &common.CaptchaCheckResp{}, errorx.CaptchaCodeError
	}
	return &common.CaptchaCheckResp{}, nil
}
