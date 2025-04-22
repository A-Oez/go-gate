package main

import (
	"go-gate/internal/db"
	sr "go-gate/internal/server"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading environment variables %v", err)
	}
	
	db, err := db.GetConnection()
	if err != nil {
		log.Println("Error getting database connection:",err)
	}
	defer db.Close()

	sr.NewServer(db).Start(":" + os.Getenv("BACKEND_PORT"))
}
