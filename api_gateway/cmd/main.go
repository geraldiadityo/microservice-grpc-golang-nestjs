package main

import (
	"api_gateway/internal/wire"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app, err := wire.InitializeServer()
	if err != nil {
		log.Fatalf("Failed to Initialize serve: %v", err)
	}

	app.Server.RegisterCleanup(func() {
		log.Println("Clean up recource")
		app.CleanupFunc()
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		if err := app.Server.Shutdown(); err != nil {
			log.Printf("Shutdown error: %v", err)
		}
	}()

	if err := app.Server.Start(os.Getenv("HTTP_PORT")); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
