package kafka_client

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	Writer *kafka.Writer
}

func NewProducer(config *Config, pCfg *ProducerConfig) *Producer {
	return &Producer{
		Writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers:      config.Brokers,
			Topic:        config.Topic,
			Balancer:     &kafka.LeastBytes{},
			BatchSize:    pCfg.BatchSize,
			BatchTimeout: time.Duration(pCfg.BatchTimeout) * time.Millisecond,
			Async:        pCfg.Async,
			RequiredAcks: pCfg.RequiredAcks,
		}),
	}
}
func (p *Producer) WriteMessages(ctx context.Context, msgs ...kafka.Message) error {
	return p.Writer.WriteMessages(ctx, msgs...)
}

// Close closes the producer writer.
func (p *Producer) Close() error {
	return p.Writer.Close()
}
