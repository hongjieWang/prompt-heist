package handlers

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/prompt-heist/backend/internal/ai"
	"github.com/prompt-heist/backend/internal/bindings"
	"github.com/prompt-heist/backend/internal/crypto"
)

type GameHandler struct {
	aiService     *ai.AIService
	signerService *crypto.SignerService
	vault         *bindings.PromptVault
}

func NewGameHandler(aiSvc *ai.AIService, signerSvc *crypto.SignerService, vault *bindings.PromptVault) *GameHandler {
	return &GameHandler{
		aiService:     aiSvc,
		signerService: signerSvc,
		vault:         vault,
	}
}

type AttemptRequest struct {
	Prompt  string `json:"prompt" binding:"required,max=2000"`
	Address string `json:"address" binding:"required,len=42"`
}

type AttemptResponse struct {
	Success   bool   `json:"success"`
	Reply     string `json:"reply"`
	Signature string `json:"signature,omitempty"`
}

func (h *GameHandler) HandleAttempt(c *gin.Context) {
	var req AttemptRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 1. Process AI Attempt
	result, err := h.aiService.ProcessAttempt(ctx, req.Prompt)
	if err != nil {
		fmt.Printf("AI Error: %v\n", err) // Add logging
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AI processing failed: " + err.Error()})
		return
	}

	fmt.Printf("AI Result: Success=%v, Reply=%s\n", result.Success, result.Reply)

	response := AttemptResponse{
		Success: result.Success,
		Reply:   result.Reply,
	}

	// 2. If success, generate signature
	fmt.Printf("response.Success=%v", response.Success)
	if response.Success {
		var amount *big.Int
		var nonce *big.Int

		if h.vault != nil {
			address := common.HexToAddress(req.Address)
			opts := &bind.CallOpts{Context: ctx}

			n, err := h.vault.Nonces(opts, address)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get nonce from contract: " + err.Error()})
				return
			}
			nonce = n

			p, err := h.vault.PrizePool(opts)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get prize pool from contract: " + err.Error()})
				return
			}
			amount = p
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Vault contract connection not initialized"})
			return
		}
		fmt.Printf("amount=%v nonce=%v", amount, nonce)

		sig, err := h.signerService.SignClaim(req.Address, amount, nonce)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Signing failed: " + err.Error()})
			return
		}
		response.Signature = sig
	}

	c.JSON(http.StatusOK, response)
}
