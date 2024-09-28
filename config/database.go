package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=" + GetEnv("DB_HOST") +
		" user=" + GetEnv("DB_USER") +
		" password=" + GetEnv("DB_PASSWORD") +
		" dbname=" + GetEnv("DB_NAME") +
		" port=" + GetEnv("DB_PORT") +
		" sslmode=disable TimeZone=Asia/Ho_Chi_Minh"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to the database successfully!")
}
