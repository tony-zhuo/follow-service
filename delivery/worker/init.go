package worker

import (
	"context"
	"github.com/tony-zhuo/follow-service/delivery/worker/init"
	"github.com/tony-zhuo/follow-service/delivery/worker/model"
	"github.com/tony-zhuo/follow-service/delivery/worker/worker_manager/follow"
	"github.com/tony-zhuo/follow-service/pkg/db"
	"github.com/tony-zhuo/follow-service/pkg/redis"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	workerMgr = make(map[string]model.WorkerManagerInterface)
	wg        *sync.WaitGroup
)

func workerInit() {
	wg = &sync.WaitGroup{}
	ctx := withContextFunc(context.Background(), func() {
		for _, worker := range workerMgr {
			worker.Shutdown()
		}
	})

	conf := init.InitConf()
	db.Init(conf.DB)
	redis.Init(conf.Redis)

	Register(follow.FollowConsumerEnable(ctx, conf.App.FollowConsumer))
}

func Run() {
	workerInit()
	enableWorker()
}

func enableWorker() {
	for _, worker := range workerMgr {
		wg.Add(1)
		go func(worker model.WorkerManagerInterface) {
			defer wg.Done()
			worker.Run()
		}(worker)
	}
	wg.Wait()
}

func Register(worker model.WorkerManagerInterface) {
	if worker == nil {
		return
	}
	workerMgr[worker.Name()] = worker
}

func withContextFunc(ctx context.Context, f func()) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal)
	go func() {
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(c)
		select {
		case <-ctx.Done():
		case <-c:
			cancel()
			f()
		}
	}()

	return ctx
}
