package main

import (
	// "api"
	// "alpaca"

	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	env := os.Getenv("BVB_ENV")
	if "" == env {
		env = "dev"
	}

	godotenv.Load(".env." + env)
}

func main() {
	loadEnv()
	port := os.Getenv("PORT")
	fmt.Printf("listening on port: %s...", port)

	serverMux := http.NewServeMux()
	// api.initialize(serverMux)
	log.Fatal(http.ListenAndServe(":"+port, serverMux))
}
