package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tigerappsorg/junction-engine/auth"
	"github.com/tigerappsorg/junction-engine/config"
	"github.com/tigerappsorg/junction-engine/database"
)

type AuthHandler struct {
	casService *auth.CASService
	db         *database.Neo4jDB
	config     *config.Config
}

func NewAuthHandler(casService *auth.CASService, db *database.Neo4jDB, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		casService: casService,
		db:         db,
		config:     cfg,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	c.Redirect(http.StatusFound, h.casService.GetLoginURL())
}

func (h *AuthHandler) Callback(c *gin.Context) {
	ticket := c.Query("ticket")
	if ticket == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No ticket provided",
		})
		return
	}

	// Validate ticket with CAS
	user, err := h.casService.ValidateTicket(ticket)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Authentication failed",
			"details": err.Error(),
		})
		return
	}

	// Store user in database (optional)
	ctx := context.Background()
	if err := h.db.CreateUser(ctx, user); err != nil {
		log.Printf("Failed to store user in database: %v", err)
		// Continue anyway - this is not critical
	}

	// Generate JWT
	token, err := h.casService.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":      token,
		"user":       user,
		"expires_in": int(h.config.JWTExpiry.Seconds()),
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":        "Logged out successfully",
		"cas_logout_url": h.casService.GetLogoutURL(),
	})
}

func (h *AuthHandler) Profile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found in context",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
