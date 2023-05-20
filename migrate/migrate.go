package main

import (
	"fmt"
	"log"

	postModel "go-rest-postgres/domain/posts/models"
	userModel "go-rest-postgres/domain/users/models"
	"go-rest-postgres/initializers"
)

func init() {
	config, err := initializers.LoadConfig()
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	initializers.DB.AutoMigrate(&userModel.User{}, &postModel.Post{})
	fmt.Println("👍 Migration complete")
}
