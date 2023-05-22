package config

import (
	"time"

	"github.com/segmentio/kafka-go"
)

var Reader *kafka.Reader

func Configure(topic string, groupId string) (kafkaReader *kafka.Reader, err error) {

	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
	}

	kafkaBrokerUrls := []string{"localhost:9092"}
	config := kafka.ReaderConfig{
		Brokers:			kafkaBrokerUrls,
		Topic:				topic,
		Dialer:				dialer,
		GroupID:			groupId,
		MinBytes:       	10e3,            // 10KB
		MaxBytes:       	10e6,            // 10MB
		MaxWait:        	1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: 	-1,
	}
	kafkaReader = kafka.NewReader(config)
	return kafkaReader, nil
}