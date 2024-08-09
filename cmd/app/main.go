package main

import (
	"blogs/api/router/v1"
	policiescasbin "blogs/internal/policies"
	"fmt"
	"log"
	"os"
	"path/filepath"

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
	// e := policiescasbin.InitEnforcer(modelPath, policyPath)

	// go func() {
	// 	ticker := time.NewTicker(10 * time.Second) // Định kỳ mỗi 5 phút
	// 	defer ticker.Stop()
	// 	for {
	// 		<-ticker.C
	// 		policiescasbin.SyncPermissions(enforcer, db, "v1")
	// 	}
	// }()
	// policiescasbin.SyncPermissions(enforcer, db, "v1")
	fmt.Println("Connect successfully!", db)

	r := router.NewRouter(db)
	r.Run(port) // listen and serve (for windows "localhost:3000")
}
