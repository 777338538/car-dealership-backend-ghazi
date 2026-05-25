package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if it exists (development convenience).
	// In production, set environment variables directly on the server.
	if err := godotenv.Load(); err != nil {
		log.Println("[CONFIG] No .env file found — using system environment variables")
	}

	// Validate critical environment variables on startup
	validateEnv()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatalf("[FATAL] Could not connect to database: %v", err)
	}

	// Seed admin account from environment variables
	store.SeedAdmin()

	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = ":2020"
	}

	server := NewAPIServer(listenAddr, store)
	server.Run()
}

// validateEnv checks that all required environment variables are set
// and logs warnings for any missing security-critical ones.
func validateEnv() {
	required := []string{"DB_PASSWORD", "JWT_SECRET", "ADMIN_EMAIL", "ADMIN_PASSWORD"}
	missing := []string{}

	for _, key := range required {
		if os.Getenv(key) == "" {
			missing = append(missing, key)
		}
	}

	if len(missing) > 0 {
		log.Printf("[SECURITY WARNING] Missing environment variables: %v", missing)
		log.Println("[SECURITY WARNING] These MUST be set before deploying to production!")
	}

	// Validate JWT secret strength
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret != "" && len(jwtSecret) < 32 {
		log.Fatal("[SECURITY] JWT_SECRET must be at least 32 characters!")
	}

	// Validate admin password strength
	adminPass := os.Getenv("ADMIN_PASSWORD")
	if adminPass != "" && len(adminPass) < 12 {
		log.Fatal("[SECURITY] ADMIN_PASSWORD must be at least 12 characters!")
	}
}
