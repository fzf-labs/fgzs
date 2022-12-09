package logic

import (
	"context"
	"fgzs/internal/app/common/model"
	"fgzs/internal/app/common/rpc/commonpb"
	"fgzs/internal/app/common/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SmsRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSmsRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SmsRecordLogic {
	return &SmsRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 短信记录
func (l *SmsRecordLogic) SmsRecord(in *commonpb.SmsRecordReq) (*commonpb.SmsRecordResp, error) {
	_, err := l.svcCtx.SmsRecordModel.Insert(l.ctx, &model.SmsRecord{
		Platform:   in.Platform,
		Phone:      in.Phone,
		SmsType:    in.SmsType,
		TmpId:      in.TmpID,
		TmpContent: in.TmpContent,
	})
	if err != nil {
		return nil, err
	}
	return &commonpb.SmsRecordResp{}, nil
}
