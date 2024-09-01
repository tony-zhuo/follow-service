package kafka

import "github.com/segmentio/kafka-go"

type Config struct {
	KafkaURL      []string `json:"url"`
	Topic         string   `json:"topic"`
	GroupId       string   `json:"group_id"`
	Offset        int64    `json:"offset"`
	MinBytes      int      `json:"min_bytes"`
	MaxBytes      int      `json:"max_bytes"`
	QueueCapacity int      `json:"queue_capacity"`
	SASLEnable    bool     `json:"sasl_enable"`
	UserName      string   `json:"user_name"`
	Password      string   `json:"password"`
}

type Option func(*Config)

var (
	OffsetLastOption Option = func(config *Config) {
		config.Offset = kafka.LastOffset
	}
	OffsetFirstOption Option = func(config *Config) {
		config.Offset = kafka.FirstOffset
	}
)

type MsgOption func(*kafka.Message)

func NewConfig(url []string, topic, groupId string, saslEnable bool, saslUser, saslPassword string) *Config {
	return &Config{
		KafkaURL:      url,
		Topic:         topic,
		GroupId:       groupId,
		QueueCapacity: 1,
		MinBytes:      1,
		MaxBytes:      10e6,
		SASLEnable:    saslEnable,
		UserName:      saslUser,
		Password:      saslPassword,
	}
}
