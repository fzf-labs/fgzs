package mq

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
	"sync"
	"time"
)

type AsynqConfig struct {
	Addr     string
	Password string
}

type Asynq struct {
	Cfg       *AsynqConfig               //配置
	Consumers map[*BusinessConfig]Handle //消费者
	Lock      sync.Mutex                 //锁
}

func NewAsynq(Cfg *AsynqConfig) *Asynq {
	return &Asynq{
		Cfg:       Cfg,
		Consumers: make(map[*BusinessConfig]Handle),
		Lock:      sync.Mutex{},
	}
}

func (a *Asynq) Publish(b *BusinessConfig, msg string) error {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: a.Cfg.Addr, Password: a.Cfg.Password})
	defer client.Close()
	_, err := client.Enqueue(asynq.NewTask(b.Topic, []byte(msg)))
	if err != nil {
		log.Printf("Asynq 延迟消息推送失败 %v", err)
		return err
	}
	return nil
}

func (a *Asynq) DeferredPublish(b *BusinessConfig, msg string, t time.Duration) error {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: a.Cfg.Addr, Password: a.Cfg.Password})
	defer client.Close()
	_, err := client.Enqueue(asynq.NewTask(b.Topic, []byte(msg)), asynq.ProcessIn(t))
	if err != nil {
		log.Printf("Asynq 延迟消息推送失败 %v", err)
		return err
	}
	return nil
}

func (a *Asynq) Register(b *BusinessConfig, handle Handle) {
	a.Lock.Lock()
	defer a.Lock.Unlock()
	a.Consumers[b] = handle
}

func (a *Asynq) Listen() {
	if len(a.Consumers) == 0 {
		fmt.Printf("Asynq 消费者注册数量为0,请核查!!!")
	}
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: a.Cfg.Addr, Password: a.Cfg.Password},
		asynq.Config{
			// 指定要使用多少并发工作人员
			Concurrency: 10,
			// 可以选择指定具有不同优先级的多个队列。
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)
	mux := asynq.NewServeMux()
	for business, handle := range a.Consumers {
		b := *business
		h := handle
		mux.HandleFunc(b.Topic, func(ctx context.Context, task *asynq.Task) error {
			err := h(string(task.Payload()))
			if err != nil {
				log.Printf("Asynq 消息业务处理失败 topic:%v,tag:%v,group_id:%v,body:%v, err:%v", b.Topic, b.Tag, b.GroupId, string(task.Payload()), err)
				return err
			}
			return nil
		})
	}
	if err := srv.Run(mux); err != nil {
		fmt.Printf("Asynq 服务启动失败: %v", err)
	}
	log.Printf("Asynq消费者监听成功,共%d个消费者", len(a.Consumers))
}
