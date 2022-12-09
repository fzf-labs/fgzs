package logic

import (
	"context"
	"fgzs/internal/app/common/rpc/common"
	"fgzs/internal/app/common/rpc/commonpb"
	"fgzs/internal/app/common/rpc/internal/svc"
	"github.com/mojocn/base64Captcha"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaGenerateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCaptchaGenerateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaGenerateLogic {
	return &CaptchaGenerateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证码
func (l *CaptchaGenerateLogic) CaptchaGenerate(in *commonpb.CaptchaGenerateReq) (*commonpb.CaptchaGenerateResp, error) {
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	captchaId, picPath, err := cp.Generate()
	if err != nil {
		return nil, err
	}
	return &common.CaptchaGenerateResp{
		CaptchaId: captchaId,
		PicPath:   picPath,
	}, nil
}
