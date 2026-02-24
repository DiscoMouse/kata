package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DiscoMouse/kata/internal/auth"
	"github.com/DiscoMouse/kata/internal/db"
	"github.com/DiscoMouse/kata/internal/manager"
)

func main() {
	// 1. Initialize dependencies (The "Parts")
	database := db.New(db.Config{URL: "postgres://..."})
	authService := auth.New(database) // Auth depends on DB

	// 2. Initialize the Manager (The "Engine")
	// You can pass the components you want to manage here
	m := manager.New(database, authService)

	// 3. Start the system in a background goroutine
	// This prevents the app from blocking while starting up
	go func() {
		log.Println("Starting system components...")
		if err := m.StartAll(); err != nil {
			log.Fatalf("Failed to start system: %v", err)
		}
		log.Println("System is RUNNING")
	}()

	// 4. Listen for shutdown signals (SIGINT, SIGTERM)
	// This ensures our "Stopping" and "Stopped" states are triggered correctly
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop // Block here until a signal is received

	// 5. Graceful Shutdown
	log.Println("Shutdown signal received. Transitioning to STOPPING...")

	// Create a timeout context so shutdown doesn't hang forever
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := m.StopAll(ctx); err != nil {
		log.Printf("Error during graceful shutdown: %v", err)
		os.Exit(1)
	}

	log.Println("System STOPPED cleanly.")
}
