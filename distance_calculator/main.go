package main

import (
	"log"

	"github.com/hamedblue1381/tolling/aggregator/client"
)

const (
	kafkaTopic         = "obudata"
	aggregatorEndpoint = "http://127.0.0.1:3000/aggregate"
)

func main() {
	var (
		err error
		svc CalculatorServicer
	)
	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)
	client := client.NewClient(aggregatorEndpoint)
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, client)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
