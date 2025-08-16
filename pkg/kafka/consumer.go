package kafka_client

import (
	"context"
	"fmt"

	"scs-operator/internal/processor"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	Reader    *kafka.Reader
	Processor *processor.Processor
}

func NewConsumer(cfg *Config, cCfg *ConsumerConfig, processor *processor.Processor) *Consumer {
	return &Consumer{
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: cfg.Brokers,
			Topic:   cfg.Topic,
			GroupID: cCfg.GroupID,
			// MinBytes:    cCfg.MinBytes,
			// MaxBytes:    cCfg.MaxBytes,
			// StartOffset: cCfg.StartOffset,
			// MaxWait:     time.Duration(cCfg.CommitInterval) * time.Millisecond,
			// Partition:   cCfg.Partition,
		}),
		Processor: processor,
	}
}
func (c *Consumer) ReadMessage(ctx context.Context) (kafka.Message, error) {
	fmt.Println("Reading message")
	return c.Reader.ReadMessage(ctx)
}

func (c *Consumer) Close() error {
	return c.Reader.Close()
}
