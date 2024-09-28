package main

import (
	"log"
	"postgres_go/Models"
	"postgres_go/Routes"
	"postgres_go/config"
)

func main() {
	config.LoadEnv()

	config.ConnectDatabase()

	config.ConnectRedis()

	err := config.DB.AutoMigrate(&Models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	router := Routes.SetupRouter()

	log.Println("Server is running on port 8080")
	router.Run(":8080")
}
