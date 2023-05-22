package main

import (
	"context"
	"fmt"
	"log"

	"go-rest-postgres/initializers"
	"go-rest-postgres/kafka/config"

	"github.com/gin-contrib/cors"
)

func main() {
	config, err := initializers.LoadConfig()
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	KafkaConsumer()
}

func KafkaConsumer() {

	kafkaReader, err := config.Configure("quickstart-events", "kafkaConsumerGroupId")
	if err != nil {
		log.Fatalln("error", err.Error())
		return
	}
	defer kafkaReader.Close()

	
	for {
		m, err := kafkaReader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("error 1 while receiving message: %s", err.Error())
			continue
		}

		value := m.Value

		if err != nil {
			fmt.Printf("error 2 while receiving message: %s", err.Error())
			continue
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, string(value))
	}

}

