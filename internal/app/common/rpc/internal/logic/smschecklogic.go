package logic

import (
	"context"
	"fgzs/internal/app/common/rpc/common"
	"fgzs/internal/app/common/rpc/commonpb"
	"fgzs/internal/app/common/rpc/internal/svc"
	"fgzs/internal/define/cachekey"
	"fgzs/internal/errorx"
	"fgzs/pkg/util/validutil"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/zeromicro/go-zero/core/logx"
)

type SmsCheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSmsCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SmsCheckLogic {
	return &SmsCheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 短信发送校验
func (l *SmsCheckLogic) SmsCheck(in *commonpb.SmsCheckReq) (*commonpb.SmsCheckResp, error) {
	//校验手机号
	if isMobile := validutil.IsPhoneLoose(in.Phone); !isMobile {
		return nil, errorx.EnterTheCorrectPhoneNumber // 请填写正确手机号
	}
	smsCacheKey := cachekey.Sms.BuildCacheKey(in.Type, in.Phone)
	code, err := l.svcCtx.Redis.GetCtx(l.ctx, smsCacheKey)
	switch err {
	case nil:
		if code != in.Code {
			return nil, errorx.SmsCodeInvalid
		}
		_, err := l.svcCtx.Redis.DelCtx(l.ctx, smsCacheKey)
		if err != nil {
			return nil, err
		}
		return &common.SmsCheckResp{}, nil
	case redis.Nil:
		return nil, errorx.SmsCodeExpired
	default:
		return nil, err
	}
}
