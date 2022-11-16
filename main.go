package main

import (
	"context"
	"runtime"

	"github.com/a1exander256/kafka/consumer"
	"github.com/a1exander256/kafka/producer"
)

func main() {
	ctx := context.Background()
	go producer.NewKafkaProducer().Producer(ctx)
	go consumer.NewKafkaReader().Cunsumer(ctx)

	runtime.Goexit()
}
