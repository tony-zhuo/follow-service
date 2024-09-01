package follow

import (
	"context"
	"fmt"
	"github.com/tony-zhuo/follow-service/delivery/worker/init"
	"github.com/tony-zhuo/follow-service/delivery/worker/model"
	"github.com/tony-zhuo/follow-service/delivery/worker/worker_manager"
)

func FollowConsumerEnable(ctx context.Context, conf *init.FollowConsumerConf) model.WorkerManagerInterface {
	if conf == nil || !conf.Enable {
		return nil
	}

	name := fmt.Sprintf("%s_%s", conf.Name, "item")
	workerManager, _ := worker_manager.NewWorkerManger(ctx, name, 1)
	workerManager.SetWorker(NewFollowWorker(ctx, conf))
	return workerManager
}
