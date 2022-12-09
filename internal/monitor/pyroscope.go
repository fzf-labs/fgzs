package monitor

import (
	"github.com/pyroscope-io/client/pyroscope"
	"github.com/zeromicro/go-zero/core/logx"
)

type PyroscopeConfig struct {
	ApplicationName string
	Addr            string
	AuthToken       string
}

type pyroscopeServer struct {
	cfg *PyroscopeConfig
}

func NewPyroscopeServer(cfg *PyroscopeConfig) *pyroscopeServer {
	return &pyroscopeServer{cfg: cfg}
}

func (p *pyroscopeServer) Start() {
	if len(p.cfg.Addr) == 0 {
		logx.Info("pyroscope server not set \n")
		return
	}
	logx.Infof("Start pyroscope server, listen addr %s\n", p.cfg.Addr)
	// 仅当您使用互斥锁或块分析时才需要这两行
	// 请阅读以下说明，了解如何设置这些费率：
	//runtime.SetMutexProfileFraction(5)
	//runtime.SetBlockProfileRate(5)
	_, err := pyroscope.Start(pyroscope.Config{
		//simple.golang.app
		ApplicationName: p.cfg.ApplicationName,

		// 将其替换为 pyroscope 服务器的地址
		// http://pyroscope-server:4040"
		ServerAddress: p.cfg.Addr,

		// 您可以通过将其设置为 nil 来禁用日志记录 pyroscope.StandardLogger
		Logger: nil,

		// 可选地，如果启用了身份验证，请指定 API 密钥：
		// AuthToken: os.Getenv("PYROSCOPE_AUTH_TOKEN"),
		AuthToken: p.cfg.AuthToken,
		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})
	if err != nil {
		logx.Error("Pyroscope start fail:", err)
		return
	}
	logx.Error("Pyroscope start success")
}

func (p *pyroscopeServer) Stop() {
	logx.Info("Pyroscope Stop")
}
