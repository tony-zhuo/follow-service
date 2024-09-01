package kafka

import (
	"github.com/segmentio/kafka-go"
	"time"
)

type KafkaTopic string

func NewKafkaWriter(config *Config, options ...Option) *kafka.Writer {
	for _, option := range options {
		option(config)
	}
	w := &kafka.Writer{
		Addr:  kafka.TCP(config.KafkaURL...),
		Topic: config.Topic,
	}

	return w
}

func GetKafkaReader(config *Config, options ...Option) *kafka.Reader {
	brokers := config.KafkaURL
	for _, option := range options {
		option(config)
	}

	cfg := kafka.ReaderConfig{
		Brokers:                brokers,
		Topic:                  config.Topic,
		GroupID:                config.GroupId,
		MinBytes:               config.MinBytes, // 10KB
		QueueCapacity:          config.QueueCapacity,
		PartitionWatchInterval: time.Duration(100 * time.Millisecond),
		WatchPartitionChanges:  true,
		ReadBackoffMin:         time.Duration(100 * time.Millisecond),
		ReadBackoffMax:         time.Duration(100 * time.Millisecond),
		MaxBytes:               config.MaxBytes, // 10MB
	}

	r := kafka.NewReader(cfg)
	r.SetOffset(config.Offset)
	return r
}

func NewMsg(topic string, key, value []byte, options ...MsgOption) kafka.Message {
	msg := &kafka.Message{
		Topic: topic,
		Key:   key,
		Value: value,
	}
	for _, option := range options {
		option(msg)
	}
	return *msg
}
