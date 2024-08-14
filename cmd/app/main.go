package main

import (
	"blogs/api/router/v1"
	policiescasbin "blogs/internal/policies"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

func redisConnection() *redis.Client {
	url := os.Getenv("REDIS_URL")
	if url == "" {
		url = "redis://localhost:6379"
	}
	opts, err := redis.ParseURL(url)
	if err != nil {
		log.Fatalf("Error parsing Redis URL: %v", err)
	}

	client := redis.NewClient(opts)
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	fmt.Println("Connected successfully to Redis!")
	return client
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DB_URL")
	port := os.Getenv("PORT")

	if dsn == "" {
		log.Fatal("Environment variable DB_URL is not set")
	}

	rdb := redisConnection()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	// // Define model and policy paths
	modelPath := filepath.Join(cwd, "config", "model.conf")
	// policyPath := filepath.Join(cwd, "config", "permissions.csv")
	_, err = policiescasbin.InitEnforcer(db, modelPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect successfully to mysql!", db)

	r := router.NewRouter(db, rdb)
	if err := r.Run(port); err != nil {
		return
	} // listen and serve (for windows "localhost:3000")
}
