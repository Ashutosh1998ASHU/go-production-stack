package main

import (
	"context"
	"fmt"
	"go-production-stack/internal/config"
	"go-production-stack/internal/handlers"
	"go-production-stack/internal/middleware"
	"go-production-stack/internal/redis"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Load Config
	cfg := config.Load()
	// Initialize Redis
	redis.InitRedis()
	// Create Router
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/live", handlers.LiveHandler)
	mux.HandleFunc("/ready", handlers.ReadyHandler)
	// Apply Middlerware
	handler := middleware.Logger(
		middleware.Recovery(mux),
	)
	// Create HTTP Server
	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: handler,
	}
	fmt.Printf("Starting server on :%s\n", server.Addr)
	// Start HTTP Server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Listen for Shutdown Signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	fmt.Println("\nShutdown Signal Recieved...")
	// Shutdown Timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Shutdown HTTP Server
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}
	// Close Redis Connection
	redis.CloseRedis()
	fmt.Println("Server Stopped Gracefully")
}
