package main

import (
	"go-gate/internal/server/router"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
    router.RegisterRouter(mux)

	log.Println("HTTPS Server runs on port :3030")

	//start server with ssl/tls
	log.Fatal(http.ListenAndServe(":3030", mux))
}
