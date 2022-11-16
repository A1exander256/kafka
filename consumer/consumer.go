package consumer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type KafkaReader struct {
	Reader *kafka.Reader
}

func NewKafkaReader() *KafkaReader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test",
	})
	return &KafkaReader{
		Reader: reader,
	}
}
func (r *KafkaReader) Cunsumer(ctx context.Context) {
	for {
		msg, _ := r.Reader.ReadMessage(ctx)

		var result map[string]interface{}
		json.Unmarshal(msg.Value, &result)
		fmt.Println("received :", result)
	}
}
