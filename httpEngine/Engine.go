package httpEngine

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
	EngineHttpServer *http.Server
	EngineContext    context.Context
}

func newEngineHttpServer() *http.Server {
	return &http.Server{
		Addr: ":8888",
	}
}

func newEngineWithConfig() *Engine {
	return &Engine{
		EngineHttpServer: newEngineHttpServer(),
		EngineContext:    context.Background(),
	}
}

/*
 *  初始化ETCD,获取ETCD配置数据
 *  初始化线程池
 *  初始后台配置
 */
func (engine *Engine) EngineWithETCD() {

}

func ServerEngine() {
	engine := newEngineWithConfig()
	engine.EngineWithETCD()

	// 启动 pprof
	go func() {
		log.Println("[pprof] starting on :6060 ...")
		if err := http.ListenAndServe("0.0.0.0:6060", nil); err != nil {
			log.Println("[pprof] failed:", err)
		}
	}()

	muxHttp := http.NewServeMux()
	muxHttp.Handle("/dsp", GzipResHandler(http.HandlerFunc(BidRequestManager)))
	engine.EngineHttpServer.Handler = muxHttp
	engine.EngineHttpServer.ListenAndServe()
	fmt.Println("停止")

}
