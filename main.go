package main

import (
	"github.com/MsN-12/url_shortener/handlers"
	"github.com/MsN-12/url_shortener/store"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	httpAddr := os.Getenv("HTTP_ADDR")
	store.InitializeStore(redisAddr, redisPassword)
	runGinServer(httpAddr)
}
func runGinServer(httpAddr string) {
	server, err := handlers.NewServer()
	if err != nil {
		log.Fatalln("Error creating server: ", err)
	}
	err = server.Start(httpAddr)
	if err != nil {
		log.Fatalln("Error starting server: ", err)
	}
}
