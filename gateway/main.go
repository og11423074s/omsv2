package main

import (
	"github.com/joho/godotenv"
	"github.com/og11423074s/omsv2/common"
	"log"
	"net/http"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":3000")
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Printf("Starting http server on %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}

}
