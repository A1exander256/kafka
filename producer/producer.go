package producer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	Writer *kafka.Writer
}

func NewKafkaProducer() *KafkaProducer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		//для докера kafka:29092 для local localhost:9092
		Brokers: []string{"kafka:29092"},
		Topic:   "test",
	})

	return &KafkaProducer{
		Writer: writer,
	}
}

func (p *KafkaProducer) Producer(ctx context.Context) {
	count := 0
	for {

		msg := map[string]interface{}{
			"key":   count,
			"value": "hello",
		}

		byteMsg, _ := json.Marshal(msg)
		p.Writer.WriteMessages(ctx, kafka.Message{
			Value: []byte(byteMsg),
		})
		count++
		time.Sleep(time.Second * 3)
	}
}
