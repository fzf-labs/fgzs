package mq

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"github.com/pkg/errors"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type NSQ struct {
	Cfg       *NsqConfig                 //配置
	Consumers map[*BusinessConfig]Handle //消费者
	Lock      sync.Mutex                 //锁
	p         []*nsq.Producer
}

func NewNSQ(cfg *NsqConfig) *NSQ {
	return &NSQ{
		Cfg:       cfg,
		Consumers: make(map[*BusinessConfig]Handle),
		Lock:      sync.Mutex{},
		p:         make([]*nsq.Producer, 0),
	}
}
func (n *NSQ) newConfig() *nsq.Config {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second
	return cfg
}

func (n *NSQ) GetProducer() (*nsq.Producer, error) {
	n.Lock.Lock()
	defer n.Lock.Unlock()
L:
	if len(n.p) == 0 {
		err := n.newProducers()
		if err != nil {
			return nil, err
		}
	}
	p := n.p[rand.Intn(len(n.p))]
	err := p.Ping()
	if err != nil {
		log.Printf("get nsq ping err:%v", err)
		p.Stop()
		//置空操作
		n.p = nil
		goto L
	}
	return p, nil
}

func (n *NSQ) newProducers() error {
	nodes := n.lookupToNodes()
	if len(nodes) == 0 {
		return errors.New("no nsq node")
	}
	n.p = make([]*nsq.Producer, 0)
	for _, node := range nodes {
		p := n.newProducer(node)
		if p != nil {
			n.p = append(n.p, p)
		}
	}
	return nil
}

func (n *NSQ) newProducer(endpoint string) *nsq.Producer {
	p, err := nsq.NewProducer(endpoint, n.newConfig())
	if err != nil {
		log.Printf("nsq NewProducer err:%v", err)
	}
	err = p.Ping()
	if err != nil {
		p.Stop()
		p = nil
		log.Printf("nsq ping err:%v", err)
	}
	log.Printf("nsq new producer success %s", endpoint)
	return p
}

func (n *NSQ) Publish(b *BusinessConfig, msg string) error {
	p, err := n.GetProducer()
	if err != nil {
		return err
	}
	return p.Publish(n.GetTopic(b), []byte(msg))
}

func (n *NSQ) DeferredPublish(b *BusinessConfig, msg string, t time.Duration) error {
	p, err := n.GetProducer()
	if err != nil {
		return err
	}
	return p.DeferredPublish(n.GetTopic(b), t, []byte(msg))
}

func (n *NSQ) Register(b *BusinessConfig, handle Handle) {
	n.Lock.Lock()
	defer n.Lock.Unlock()
	n.Consumers[b] = handle
}

func (n *NSQ) Listen() {
	if len(n.Consumers) > 0 {
		for business, handle := range n.Consumers {
			b := *business
			h := handle
			go n.do(&b, h)
		}
	}
	log.Printf("nsq消费者监听成功,共%d个消费者", len(n.Consumers))
}

func (n *NSQ) do(b *BusinessConfig, handle Handle) {
	q, err := nsq.NewConsumer(n.GetTopic(b), n.GetChannel(b), n.newConfig())
	if err != nil {
		return
	}
	q.SetLoggerLevel(nsq.LogLevelError)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		err := handle(string(message.Body))
		if err != nil {
			log.Printf("nsq 消息业务处理失败 topic:%v,tag:%v,group_id:%v,body:%v, err:%v", b.Topic, b.Tag, b.GroupId, string(message.Body), err)
			return err
		}
		return nil
	}))
	_ = q.ConnectToNSQLookupds(n.Cfg.Lookupds)
}

func (n *NSQ) GetTopic(b *BusinessConfig) string {
	return b.Topic + "_" + b.Tag
}
func (n *NSQ) GetChannel(b *BusinessConfig) string {
	return b.Tag
}

// 获取nsqd的节点信息
func (n *NSQ) lookupToNodes() []string {
	lookups := n.Cfg.Lookupds
	var nodes []string
	for _, lookup := range lookups {
		lookupNodes, err := n.getLookupNodes(lookup)
		if err != nil {
			continue
		}
		nodes = lookupNodes
		break
	}
	return nodes
}

type NodesRsp struct {
	Producers []struct {
		RemoteAddress    string `json:"remote_address"`
		Hostname         string `json:"hostname"`
		BroadcastAddress string `json:"broadcast_address"`
		TcpPort          int    `json:"tcp_port"`
		HttpPort         int    `json:"http_port"`
		Version          string `json:"version"`
	} `json:"producers"`
}

func (n *NSQ) getLookupNodes(lookup string) ([]string, error) {
	nodes := make([]string, 0)
	url := "http://" + lookup + "/nodes"
	ret, err := n.httpGet(url)
	if err != nil {
		return nodes, err
	}
	var rsp NodesRsp
	err = json.Unmarshal(ret, &rsp)
	if err != nil {
		return nodes, err
	}
	if len(rsp.Producers) > 0 {
		for _, v := range rsp.Producers {
			nodes = append(nodes, v.BroadcastAddress+":"+strconv.Itoa(v.TcpPort))
		}
	}
	return nodes, nil
}

func (n *NSQ) httpGet(url string) ([]byte, error) {
	client := &http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
