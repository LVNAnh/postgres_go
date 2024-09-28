package config

import (
    "context"
    "github.com/redis/go-redis/v9"
    "log"
)

var RDB *redis.Client
var Ctx = context.Background()


func ConnectRedis() {
    RDB = redis.NewClient(&redis.Options{
        Addr:     GetEnv("REDIS_HOST") + ":" + GetEnv("REDIS_PORT"),
        Password: GetEnv("REDIS_PASSWORD"), 
        DB:       0,                       
    })

    _, err := RDB.Ping(Ctx).Result()
    if err != nil {
        log.Fatalf("Failed to connect to Redis: %v", err)
    }

    log.Println("Connected to Redis successfully!")
}
