package mq

import (
	"fmt"
)

var BusinessConfigs = map[string]*BusinessConfig{}

// BusinessConfig 不同业务的配置
type BusinessConfig struct {
	// Topic所属的实例ID，在消息队列RocketMQ版控制台创建。
	// 若实例有命名空间，则实例ID必须传入；若实例无命名空间，则实例ID传入null空值或字符串空值。实例的命名空间可以在消息队列RocketMQ版控制台的实例详情页面查看。
	Name string `json:"name"`
	// 消息所属的Topic，在消息队列RocketMQ版控制台创建。
	//不同消息类型的Topic不能混用，例如普通消息的Topic只能用于收发普通消息，不能用于收发其他类型的消息。
	Topic string `json:"topic"`
	//标签
	Tag string `json:"tag"`
	// 您在控制台创建的Group ID。
	GroupId string `json:"group_id"`
}

func NewBusiness(name string, topic string, tag string, groupId string) *BusinessConfig {
	prefixName := topic + tag + groupId
	if _, ok := BusinessConfigs[prefixName]; ok {
		panic(fmt.Sprintf("mq key %s is exsit, please change one", prefixName))
	}
	b := &BusinessConfig{Name: name, Topic: topic, Tag: tag, GroupId: groupId}
	BusinessConfigs[prefixName] = b
	return b
}

func GetBusiness(key string) (*BusinessConfig, error) {
	business, ok := BusinessConfigs[key]
	if !ok {
		return nil, KeyNotFound
	}
	return &BusinessConfig{
		Name:    business.Name,
		Topic:   business.Topic,
		Tag:     business.Tag,
		GroupId: business.GroupId,
	}, nil
}
