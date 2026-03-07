package main

import (
	"context"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/prompt-heist/backend/api"
	"github.com/prompt-heist/backend/internal/ai"
	"github.com/prompt-heist/backend/internal/bindings"
	"github.com/prompt-heist/backend/internal/crypto"
	"github.com/prompt-heist/backend/internal/handlers"
	"github.com/prompt-heist/backend/internal/listener"
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

	// 3. Setup Contract Binding for Read Calls
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		rpcURL = "https://bsc-testnet-rpc.publicnode.com" // default public HTTP RPC for BSC Testnet
	}

	contractAddress := os.Getenv("CONTRACT_ADDRESS")
	if contractAddress == "" {
		contractAddress = "0x117D20BdF529891421546dc5F8651561A0F59aE0" // Current deployed PromptVault
	}

	var vault *bindings.PromptVault
	var chainID *big.Int
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Printf("Failed to connect to Ethereum RPC node: %v", err)
	} else {
		addr := common.HexToAddress(contractAddress)
		vault, err = bindings.NewPromptVault(addr, ethClient)
		if err != nil {
			log.Printf("Failed to bind PromptVault contract: %v", err)
			vault = nil
		}
		chainID, err = ethClient.ChainID(context.Background())
		if err != nil {
			log.Printf("Failed to fetch chain ID: %v", err)
			chainID = nil
		}
	}
	vaultAddress := common.HexToAddress(contractAddress)

	// 4. Init Handlers
	gameHandler := handlers.NewGameHandler(aiService, signerService, vault, chainID, vaultAddress)

	// 5. Init Router
	router := api.NewRouter(gameHandler)

	// 6. Start Event Listener
	wsURL := os.Getenv("WS_URL")
	if wsURL == "" {
		wsURL = "wss://bsc-testnet-rpc.publicnode.com" // default public WS for BSC Testnet
	}

	eventListener, err := listener.NewEventListener(wsURL, contractAddress)
	if err != nil {
		log.Printf("Failed to initialize event listener (are you using a valid WebSocket URL?): %v", err)
	} else {
		// Start in a goroutine
		go eventListener.Start(context.Background())
	}

	// 7. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
