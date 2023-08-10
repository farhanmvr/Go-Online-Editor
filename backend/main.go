package main

import (
	"fmt"
	"github.com/farhanmvr/go-editor/config"
	"github.com/farhanmvr/go-editor/db"
	"github.com/farhanmvr/go-editor/routes"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	// Load config
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error while loading the config: %v", err)
	}
	config := config.GetConfig()

	// Configure db
	db.InitDB()            // Initialize database connection
	db.PerformMigrations() // Apply migrations

	// Create a new CORS handler with desired options
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:3333"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Setup routes
	router := routes.SetupRoutes()

	handler := c.Handler(router)

	port := config.Server.Port
	fmt.Println("Server started running successfully")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
