package main

import (
	"github.com/MsN-12/url_shortener/handlers"
	"github.com/MsN-12/url_shortener/store"
	"log"
)

func main() {
	store.InitializeStore()
	runGinServer()
}
func runGinServer() {
	server, err := handlers.NewServer()
	if err != nil {
		log.Fatalln("Error creating server: ", err)
	}
	err = server.Start(":3000")
	if err != nil {
		log.Fatalln("Error starting server: ", err)
	}
}
