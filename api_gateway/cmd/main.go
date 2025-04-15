package main

import (
	"api_gateway/internal/wire"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server, err := wire.InitializeServer()
	if err != nil {
		log.Fatalf("Failed to Initialize serve: %v", err)
	}

	server.RegisterCleanup(func() {
		log.Println("Clean up recource")
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		if err := server.Shutdown(); err != nil {
			log.Printf("Shutdown error: %v", err)
		}
	}()

	if err := server.Start(os.Getenv("HTTP_PORT")); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
