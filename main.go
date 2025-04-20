package main

import (
	"go-gate/internal/db"
	sr "go-gate/internal/server"
	"log"
	"net/http"
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

	server := sr.NewServer(db)
	server.RegisterRouter()

	port := ":" + os.Getenv("BACKEND_PORT")
	log.Printf("HTTPS Server runs on port %s", port)
	log.Fatal(http.ListenAndServe(port, server.Mux))
}
