package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/prompt-heist/backend/api"
	"github.com/prompt-heist/backend/internal/ai"
	"github.com/prompt-heist/backend/internal/crypto"
	"github.com/prompt-heist/backend/internal/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	// 1. Init Signer
	privateKeyHex := os.Getenv("SIGNER_PRIVATE_KEY")
	if privateKeyHex == "" {
		log.Fatal("SIGNER_PRIVATE_KEY is required")
	}
	signerService, err := crypto.NewSignerService(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to init signer: %v", err)
	}
	log.Printf("Signer initialized with address: %s", signerService.GetAddress())

	// 2. Init AI
	aiService := ai.NewAIService()

	// 3. Init Handlers
	gameHandler := handlers.NewGameHandler(aiService, signerService)

	// 4. Init Router
	router := api.NewRouter(gameHandler)

	// 5. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
