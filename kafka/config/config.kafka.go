package config

import (
	"flag"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

var Writer *kafka.Writer

var (
	kafkaBrokerUrl string
	kafkaVerbose   bool
	kafkaClientId  string
)

func Configure(topic string) (w *kafka.Writer, err error) {
	flag.StringVar(&kafkaBrokerUrl, "kafka-brokers", "localhost:9092", "Kafka brokers in comma separated value")
	flag.BoolVar(&kafkaVerbose, "kafka-verbose", true, "Kafka verbose logging")
	flag.StringVar(&kafkaClientId, "kafka-client-id", "my-kafka-client", "Kafka client id to connect")

	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: kafkaClientId,
	}

	var kafkaBrokerUrls []string

	kafkaBrokerUrls = strings.Split(kafkaBrokerUrl, ",")
	config := kafka.WriterConfig{
		Brokers:          kafkaBrokerUrls,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	w = kafka.NewWriter(config)
	Writer = w
	return w, nil
}