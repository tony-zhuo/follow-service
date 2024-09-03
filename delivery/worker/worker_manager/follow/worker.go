package follow

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/tony-zhuo/follow-service/delivery/worker/init"
	pkgKafka "github.com/tony-zhuo/follow-service/pkg/kafka"
	"github.com/tony-zhuo/follow-service/pkg/libs"
	"github.com/tony-zhuo/follow-service/service/model"
	"strings"
	"sync"
)

var (
	followWorker     *FollowWorker
	followWorkerOnce sync.Once
)

type FollowWorker struct {
	ctx         context.Context
	conf        *init.FollowConsumerConf
	followUc    model.FollowUcInterface
	kafkaReader *kafka.Reader
}

func NewFollowWorker(ctx context.Context, conf *init.FollowConsumerConf) *FollowWorker {
	followWorkerOnce.Do(func() {
		followWorker = &FollowWorker{
			ctx:  ctx,
			conf: conf,
		}
	})
	return followWorker
}

func (w *FollowWorker) ID(ctx context.Context) string {

	return ""
}

func (w *FollowWorker) Exec(ctx context.Context) error {
	topic := fmt.Sprintf("%s", w.conf.Topic)
	cfg := pkgKafka.NewConfig(w.conf.KafkaServerUrl, topic, w.conf.Group, w.conf.SASLEnable, w.conf.UserName, w.conf.Password)
	kafkaReader := pkgKafka.GetKafkaReader(cfg, pkgKafka.OffsetLastOption)
	defer func() {
		kafkaReader.Close()
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			message, err := kafkaReader.ReadMessage(context.Background())
			if err != nil {
				// logging
				continue
			}

			keys := strings.Split(string(message.Key), "_")
			action := model.Action(keys[0])

			timestamp, err := libs.StringConvertToTime(keys[3])
			if err != nil {
				// logging
				continue
			}

			if !w.followUc.CheckAndSyncConsumerTime(ctx, *timestamp) {
				// logging
				continue
			}

			obj := &model.FollowRequest{}
			if err := json.Unmarshal(message.Value, obj); err != nil {
				// logging
				continue
			}

			switch action {
			case model.Action_Follow:
				if err := w.followUc.FollowInDB(ctx, obj); err != nil {
					// logging
					continue
				}
			case model.Action_UnFollow:
				if err := w.followUc.UnFollowInDB(ctx, obj); err != nil {
					// logging
					continue
				}
			}
		}
	}
}

func (w *FollowWorker) Topic(ctx context.Context) string {
	return w.conf.Topic
}

func (w *FollowWorker) Group(ctx context.Context) string {
	return w.conf.Group
}

func (w *FollowWorker) Health(ctx context.Context) bool {
	return true
}
