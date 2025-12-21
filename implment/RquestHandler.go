package implment

import (
	"github.com/cxb116/DSP/constant"
	"log"
	"sync"
)

const (
	WORKER_POOL_SIZE   = 2000 // 池大小
	MAX_CHANNLE_LENGHT = 2000
)

type WorkerChannelHandler struct {
	WorkerPoolSize   int                 // worker 池子大小
	MaxWorkerTaskLen int                 // worker 管道队列长度
	FreeWorkers      map[int]struct{}    // 空闲worker集合
	TaskQueue        []chan *interface{} //Worker负责取任务的消息队列
	FreeWorkerMutex  *sync.RWMutex
}

// 初始化工作池
func NewWorkerChannelHandler(workerPoolSize int, maxWorkerTaskLen int) *WorkerChannelHandler {
	var freeWorkers map[int]struct{}
	freeWorkers = make(map[int]struct{}, workerPoolSize)
	for i := 0; i < workerPoolSize; i++ {
		freeWorkers[i] = struct{}{}
	}

	return &WorkerChannelHandler{
		WorkerPoolSize:   workerPoolSize,
		MaxWorkerTaskLen: maxWorkerTaskLen,
		FreeWorkers:      freeWorkers,
		TaskQueue:        make([]chan *interface{}, workerPoolSize),
		FreeWorkerMutex:  new(sync.RWMutex),
	}
}

// 启动工作池
func (this *WorkerChannelHandler) StartWorkerPool() {
	for i := 0; i < this.WorkerPoolSize; i++ {
		this.TaskQueue[i] = make(chan *interface{}, this.MaxWorkerTaskLen)
		go this.OnWorkerHandle(i, this.TaskQueue[i])
	}
}

// 获取workerId
func (this *WorkerChannelHandler) useWorker() int {
	var workerId int
	this.FreeWorkerMutex.Lock()
	defer this.FreeWorkerMutex.Unlock()

	for workerId = range this.FreeWorkers {
		delete(this.FreeWorkers, workerId)
		return workerId
	}
	return constant.NOT_WORKER
}

// 释放workerId
func (this *WorkerChannelHandler) releaseWorker(workerId int) {
	this.FreeWorkerMutex.Lock()
	defer this.FreeWorkerMutex.Unlock()
	this.FreeWorkers[workerId] = struct{}{}
}

func (this *WorkerChannelHandler) OnWorkerHandle(workerId int, Task chan *interface{}) {

	for {
		select {}
	}
}

func (this *WorkerChannelHandler) doRequestDispathcher() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("doRequestDispathcher err:", err)
		}
	}()
}
