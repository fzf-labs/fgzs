package server

import (
	"context"
	"fgzs/internal/app/mq/rpc/internal/logic/member"
	"fgzs/internal/app/mq/rpc/internal/svc"
	"fgzs/internal/define/mqkey"
	"fmt"
)

type MqService struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMqService(ctx context.Context, svcContext *svc.ServiceContext) *MqService {
	return &MqService{
		ctx:    ctx,
		svcCtx: svcContext,
	}
}

func (m *MqService) Start() {
	fmt.Println("mq Start begin")

	// 用户注销
	m.svcCtx.MqClient.Register(mqkey.UserCancellation, member.NewMemberCancellationLogicLogic(m.ctx, m.svcCtx).MemberCancellation)

	fmt.Println("mq Start done !")
}

func (m *MqService) Stop() {
	fmt.Println("mq Stop")
}
