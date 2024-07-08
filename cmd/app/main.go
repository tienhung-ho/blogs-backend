package main

import (
	"blogs/api/router/v1"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DB_URL")
	port := os.Getenv("PORT")

	if dsn == "" {
		log.Fatal("Environment variable DB_URL is not set")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connect successfully!", db)

	r := router.NewRouter(db)
	r.Run(port) // listen and serve (for windows "localhost:3000")

}
