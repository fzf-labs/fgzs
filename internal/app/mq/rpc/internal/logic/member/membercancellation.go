package member

import (
	"context"
	"fgzs/internal/app/mq/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type MemberCancellationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberCancellationLogicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberCancellationLogic {
	return &MemberCancellationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (m *MemberCancellationLogic) MemberCancellation(msg string) error {

	return nil
}
