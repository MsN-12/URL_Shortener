package main

import (
	"github.com/MsN-12/url_shortener/api"
	"github.com/MsN-12/url_shortener/store"
	"log"
)

func main() {
	store.InitializeStore()
	runGinServer()
}
func runGinServer() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatalln("Error creating server: ", err)
	}
	err = server.Start(":8080")
	if err != nil {
		log.Fatalln("Error starting server: ", err)
	}
}
