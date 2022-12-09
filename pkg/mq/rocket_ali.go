package mq

import (
	mq_http_sdk "github.com/aliyunmq/mq-http-go-sdk"
	"github.com/gogap/errors"
	"log"
	"strings"
	"sync"
	"time"
)

func NewRocketAli(cfg *RocketMqAliConfig) *RocketAli {
	return &RocketAli{
		Cfg:       cfg,
		Consumers: make(map[*BusinessConfig]Handle),
		Lock:      sync.Mutex{},
		once:      sync.Once{},
		mqClient:  nil,
	}
}

type RocketAli struct {
	Cfg       *RocketMqAliConfig         //配置
	Consumers map[*BusinessConfig]Handle //消费者
	Lock      sync.Mutex                 //锁
	once      sync.Once
	mqClient  mq_http_sdk.MQClient
}

func (r *RocketAli) newMQClient() mq_http_sdk.MQClient {
	r.once.Do(
		func() {
			endpoint := r.Cfg.Endpoint
			accessKey := r.Cfg.AccessKey
			secretKey := r.Cfg.SecretKey
			r.mqClient = mq_http_sdk.NewAliyunMQClient(endpoint, accessKey, secretKey, "")
		})
	return r.mqClient
}

func (r *RocketAli) Publish(b *BusinessConfig, msg string) error {
	//消息体封装
	msgReq := mq_http_sdk.PublishMessageRequest{
		MessageBody: msg,   // 消息内容
		MessageTag:  b.Tag, // 消息标签
	}
	instanceId := r.Cfg.InstanceId
	mqProducer := r.newMQClient().GetProducer(instanceId, b.Topic)
	ret, err := mqProducer.PublishMessage(msgReq)
	if err != nil {
		return GeneralMessageDeliveryFailed
	}
	log.Printf("阿里云rocketmq 普通消息生产成功,topic:%s,tag:%s,MsgID:%s,Body:%s", b.Topic, b.Tag, ret.MessageId, msg)
	return nil
}

func (r *RocketAli) DeferredPublish(b *BusinessConfig, msg string, t time.Duration) error {
	//消息体封装
	msgReq := mq_http_sdk.PublishMessageRequest{
		MessageBody:      msg,   // 消息内容
		MessageTag:       b.Tag, // 消息标签
		StartDeliverTime: time.Now().Add(t).UTC().Unix() * 1000,
	}
	instanceId := r.Cfg.InstanceId
	mqProducer := r.newMQClient().GetProducer(instanceId, b.Topic)
	ret, err := mqProducer.PublishMessage(msgReq)
	if err != nil {
		return DelayedMessageDeliveryFailed
	}
	log.Printf("阿里云rocketmq 延时消息生产成功,topic:%s,tag:%s,MessageId:%s,Body:%s", b.Topic, b.Tag, ret.MessageId, msgReq.MessageBody)
	return nil
}

func (r *RocketAli) Register(b *BusinessConfig, handle Handle) {
	r.Lock.Lock()
	defer r.Lock.Unlock()
	r.Consumers[b] = handle
}

func (r *RocketAli) Listen() {
	if len(r.Consumers) > 0 {
		for business, handle := range r.Consumers {
			b := *business
			h := handle
			go r.do(&b, h)
		}
	}
	log.Printf("阿里云rocketmq消费者监听成功,共%d个消费者", len(r.Consumers))
}

func (r *RocketAli) do(b *BusinessConfig, handle Handle) {
	instanceId := r.Cfg.InstanceId
	consumer := r.newMQClient().GetConsumer(instanceId, b.Topic, b.GroupId, b.Tag)
	for {
		endChan := make(chan int)
		respChan := make(chan mq_http_sdk.ConsumeMessageResponse)
		errChan := make(chan error)

		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Printf(b.Name+"捕获异常 err=%v", err)
					endChan <- 1
				}
			}()

			select {
			case resp := <-respChan:
				{
					log.Printf(b.Name+"收到消息 ----> msg=%+v", resp.Messages)

					// 处理业务逻辑
					var handles []string
					for _, v := range resp.Messages {
						err := handle(v.MessageBody)
						if err != nil {
							log.Printf(b.Name+"业务处理失败 msgBody=%s err=%v", v.MessageBody, err)
							continue
						}
						handles = append(handles, v.ReceiptHandle)
					}

					// NextConsumeTime前若不确认消息消费成功，则消息会重复消费
					// 消息句柄有时间戳，同一条消息每次消费拿到的都不一样
					ackerr := consumer.AckMessage(handles)
					if ackerr != nil {
						// 某些消息的句柄可能超时了会导致确认不成功
						for _, errAckItem := range ackerr.(errors.ErrCode).Context()["Detail"].([]mq_http_sdk.ErrAckItem) {
							log.Printf(b.Name+"消息确认失败 ErrorHandle:%s, ErrorCode:%s, ErrorMsg:%s",
								errAckItem.ErrorHandle, errAckItem.ErrorCode, errAckItem.ErrorMsg)
						}
						time.Sleep(3 * time.Second)
					} else {
						log.Printf(b.Name+"消费成功 ----> handles=%v", handles)
					}

					endChan <- 1
				}
			case err := <-errChan:
				{
					// 没有消息
					if strings.Contains(err.(errors.ErrCode).Error(), "MessageNotExist") {
					} else {
						log.Printf(b.Name+"接收消息失败 err=%v", err)
						time.Sleep(3 * time.Second)
					}
					endChan <- 1
				}
			case <-time.After(35 * time.Second):
				{
					log.Printf(b.Name + "消息消费超时")
					endChan <- 1
				}
			}
		}()

		// 长轮询消费消息
		// 长轮询表示如果topic没有消息则请求会在服务端挂住3s，3s内如果有消息可以消费则立即返回
		consumer.ConsumeMessage(respChan, errChan,
			3, // 一次最多消费3条(最多可设置为16条)
			3, // 长轮询时间3秒（最多可设置为30秒）
		)
		<-endChan
	}
}
