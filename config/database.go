package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// func ConnectDatabase() {
// 	dsn := "host=" + GetEnv("DB_HOST") +
// 		" user=" + GetEnv("DB_USER") +
// 		" password=" + GetEnv("DB_PASSWORD") +
// 		" dbname=" + GetEnv("DB_NAME") +
// 		" port=" + GetEnv("DB_PORT") +
// 		" sslmode=disable TimeZone=Asia/Ho_Chi_Minh"

// 	var err error
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
// 	}

// 	log.Println("Connected to the database successfully!")
// }

func ConnectDatabase() {
	dsn := GetEnv("DB_USER") + ":" + GetEnv("DB_PASSWORD") + "@tcp(127.0.0.1:3306)/" + GetEnv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to the MySQL database successfully!")
}
