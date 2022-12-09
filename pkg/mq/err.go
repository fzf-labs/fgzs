package mq

import "github.com/pkg/errors"

var (
	KeyNotFound                  = errors.New("MQ配置key未找到")
	GeneralMessageDeliveryFailed = errors.New("普通消息投递失败")
	DelayedMessageDeliveryFailed = errors.New("延时消息投递失败")
	DelayLevelError              = errors.New("rocketmq延时等级错误")
)
