package monitor

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type pprofServer struct {
	addr string
}

func NewPprofServer(addr string) *pprofServer {
	return &pprofServer{addr: addr}
}

func (p *pprofServer) Start() {
	if len(p.addr) == 0 {
		logx.Info("pyroscope server not set \n")
		return
	}
	logx.Infof("Start pprof server, listen addr %s\n", p.addr)
	err := http.ListenAndServe(p.addr, nil)
	if err != nil {
		logx.Error("Start pprof server err:", err)
	}
}

func (p *pprofServer) Stop() {
	logx.Info("Stop pprof server")
}
