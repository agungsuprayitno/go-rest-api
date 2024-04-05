package main

import (
	"fmt"
	"log"

	authModel "go-rest-postgres/domain/auth/models"
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
	initializers.DB.AutoMigrate(&authModel.Authorization{}, &authModel.Merchant{}, &authModel.MerchantHistory{})
	fmt.Println("👍 Migration complete")
}
