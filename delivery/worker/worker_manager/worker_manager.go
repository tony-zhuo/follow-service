package worker_manager

import (
	"context"
	"github.com/tony-zhuo/follow-service/delivery/worker/model"
	"sync"
)

type WorkerManager struct {
	ctx             context.Context
	cancel          context.CancelFunc
	wg              *sync.WaitGroup
	name            string
	count           int
	workerInterface model.WorkerInterface
	workers         []model.WorkerInterface
}

func NewWorkerManger(ctx context.Context, name string, routineCount int) (*WorkerManager, error) {
	c, cancel := context.WithCancel(ctx)
	tradeParserManager := &WorkerManager{
		ctx:     c,
		cancel:  cancel,
		wg:      &sync.WaitGroup{},
		name:    name,
		count:   routineCount,
		workers: make([]model.WorkerInterface, routineCount),
	}

	return tradeParserManager, nil
}

func (pm *WorkerManager) SetWorker(worker model.WorkerInterface) {
	pm.workerInterface = worker
}

func (pm *WorkerManager) Name() string {
	return pm.name
}

func (pm *WorkerManager) Run() error {
	//logs start
	for i := 0; i < pm.count; i++ {
		pm.wg.Add(1)
		pm.workers[i] = pm.workerInterface
		go func(idx int) {
			defer func() {
				pm.wg.Done()
				//logs "[workers] %d done", idx
			}()
			//logs "[workers] %d done", idx
			pm.workers[idx].Exec(pm.ctx)
		}(i)
	}
	pm.wg.Wait()
	pm.Clear()
	//logs end
	return nil
}

func (pm *WorkerManager) Clear() error {
	pm.workers = nil
	return nil
}

func (pm *WorkerManager) Shutdown() error {
	//logs cancel
	pm.cancel()
	//logs shutdown
	return nil
}

func (pm *WorkerManager) Health() bool {
	for _, worker := range pm.workers {
		if worker.Health(pm.ctx) {
			return true
		}
	}
	return false
}
