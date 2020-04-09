package main

import (
	// "api"
	// "alpaca"

	"net/http"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func loadEnv() {
	env := os.Getenv("BVB_ENV")
	if "" == env {
		env = "dev"
	}

	godotenv.Load(".env." + env)
}

// Sets the logging configurations
func setLogConfigs() *os.File {
	log.SetLevel(log.WarnLevel)
	if os.Getenv("BVB_ENV") == "dev" || os.Getenv("BVB_ENV") == "" {
		log.SetLevel(log.DebugLevel)
	}

	file, err := os.OpenFile("./log/info.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if os.Getenv("BVB_ENV") != "dev" {
		// TODO: Send to a prod/stg cloud log file.
	}

	if err != nil {
		log.Warn(err)
	}

	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{})

	return file
}

func main() {
	loadEnv()
	file := setLogConfigs()
	defer file.Close()

	port := os.Getenv("PORT")
	log.Info("listening on port: %s...", port)

	serverMux := http.NewServeMux()

	// api.initialize(serverMux)
	log.Fatal(http.ListenAndServeTLS(
		":"+port,
		"https-server.crt",
		"https-server.key",
		serverMux,
	))
}
