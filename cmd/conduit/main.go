package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/quantinium03/conduit/internal/database"
	"github.com/quantinium03/conduit/pkg/routes"
	"github.com/quantinium03/conduit/pkg/types"

    _ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get PORT from environment variables or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		log.Printf("PORT is not found in the environment variable. Defaulting to PORT 8080")
		port = "8080"
	}

	// Get DSN_URI from environment variables for the database connection
	/* dsnURI := os.Getenv("DSN_URI")
	if dsnURI == "" {
		log.Fatal("Failed to find DSN_URI in environment variable")
	} */

	// Open database connection
	/* conn, err := sql.Open("sqlite", "conduit.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer conn.Close() // Ensure that the database connection is closed when the program ends

	// Initialize the database connection object
	db := database.New(conn)
*/
	// Setup the API configuration
	/* apiCfg := types.ApiConfig{
		DB: db,
	} */

	// Create a new router
	router := chi.NewRouter()

	// CORS middleware configuration
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // You can modify this based on your needs (e.g., specific domains)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Set up versioned routes
	v1Router := chi.NewRouter()
	routes.SetupRoutes(v1Router)

	// Mount the versioned router under "/v1"
	router.Mount("/v1", v1Router)

	// Start the HTTP server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server starting at :%v", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
