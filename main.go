package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go-rest-postgres/initializers"
	"go-rest-postgres/kafka/config"
	"go-rest-postgres/kafka/producer"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server              *gin.Engine
)

func init() {
	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig()
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("")
	router.POST("/kafka-producer", KafkaMessage)

	log.Fatal(server.Run(":" + config.ServerPort))
}

func KafkaMessage(ctx *gin.Context) {

	// Parse JSON data from the request body
	var jsonMessage map[string]interface{}
	if err := ctx.ShouldBindJSON(&jsonMessage); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Convert JSON message to string
	stringMessage := fmt.Sprintf("%v", jsonMessage)

	message, err := json.Marshal(stringMessage)

	kafkaProducer, err := config.Configure("quickstart-events")
	if err != nil {
		log.Fatalln("error", err.Error())
		return
	}
	defer kafkaProducer.Close()

	err = producer.PushMessage(context.Background(), nil, message)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": map[string]interface{}{
				"message": fmt.Sprintf("error while push message into kafka: %s", err.Error()),
			},
		})

		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "success push data into kafka",
		"data":    stringMessage,
	})
}

