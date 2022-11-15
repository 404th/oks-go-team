package main

import (
	"log"

	"github.com/404th/todo/server"
)

func main() {
	srv := new(server.Server)

	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error performed while starting server: %v", err)
	}
}
