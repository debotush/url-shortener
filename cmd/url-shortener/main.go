package main

import (
	"log"
	"net/http"
	customHttp "url-shortener-service/api/http"
	customConfig "url-shortener-service/config"
	"url-shortener-service/internal/database"
)

func main() {
	// Load the configuration
	cfg := customConfig.LoadConfig()

	// Create a new HTTP server instance
	server := customHttp.NewServer(cfg)

	// Register request handlers
	server.RegisterHandlers()

	// Initialize the Gorm database
	database.InitializeGormDatabase()

	// Print a message indicating that the server is running
	log.Println("Server is running on", cfg.ServerAddress)

	// Start the server
	err := server.Start()

	// Handle server start error
	if err != nil && err == http.ErrServerClosed {
		log.Fatal("Server crash due to:", err)
	}
}
