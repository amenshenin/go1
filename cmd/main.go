package main

import (
	"log"

	todo "github.com/amenshenin/go1"
	"github.com/amenshenin/go1/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRouts()); err != nil {
		log.Fatalf("error occured while running http server %w", err.Error())
	}
}
