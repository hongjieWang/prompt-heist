package handlers

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prompt-heist/backend/internal/ai"
	"github.com/prompt-heist/backend/internal/crypto"
)

type GameHandler struct {
	aiService     *ai.AIService
	signerService *crypto.SignerService
}

func NewGameHandler(aiSvc *ai.AIService, signerSvc *crypto.SignerService) *GameHandler {
	return &GameHandler{
		aiService:     aiSvc,
		signerService: signerSvc,
	}
}

type AttemptRequest struct {
	Prompt  string `json:"prompt" binding:"required"`
	Address string `json:"address" binding:"required"`
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

	response := AttemptResponse{
		Success: result.Success,
		Reply:   result.Reply,
	}

	// 2. If success, generate signature
	if result.Success {
		// Mock logic: Sign for a winning claim.
		// In a real app, we fetch the current prize pool amount and user's nonce from chain.
		// Here, we hardcode 0 for simplicity, assuming the contract is updated to not verify amount,
		// or we update this later.
		amount := big.NewInt(0)
		nonce := big.NewInt(0)

		sig, err := h.signerService.SignClaim(req.Address, amount, nonce)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Signing failed: " + err.Error()})
			return
		}
		response.Signature = sig
	}

	c.JSON(http.StatusOK, response)
}
