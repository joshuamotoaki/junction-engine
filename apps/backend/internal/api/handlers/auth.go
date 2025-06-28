package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tigerappsorg/junction-engine/internal/database/neo4j"
	"github.com/tigerappsorg/junction-engine/internal/shared/auth"
	"github.com/tigerappsorg/junction-engine/internal/shared/config"
)

type authHandler struct {
	casService auth.CASService
	db         neo4j.Neo4jDB
	config     *config.Config
}

type AuthHandler interface {
	Login(c *gin.Context)
	Callback(c *gin.Context)
	Logout(c *gin.Context)
	Profile(c *gin.Context)
}

func NewAuthHandler(casService auth.CASService, db neo4j.Neo4jDB, cfg *config.Config) AuthHandler {
	return &authHandler{
		casService: casService,
		db:         db,
		config:     cfg,
	}
}

// @Summary		Login with CAS
// @Description	Redirects to CAS login page
// @Tags			auth
// @Accept			json
// @Produce		json
// @Success		302
// @Router			/auth/login [get]
func (h *authHandler) Login(c *gin.Context) {
	c.Redirect(http.StatusFound, h.casService.GetLoginURL())
}

// @Summary		CAS Callback
// @Description	Handles CAS callback and generates JWT token
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			ticket	query		string	true	"CAS Ticket"
// @Success		200		{object}	map[string]interface{}
// @Failure		400		{object}	map[string]string	"Bad Request"
// @Failure		401		{object}	map[string]string	"Unauthorized"
// @Failure		500		{object}	map[string]string	"Internal Server Error"
// @Router			/auth/callback [get]
func (h *authHandler) Callback(c *gin.Context) {
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

func (h *authHandler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":        "Logged out successfully",
		"cas_logout_url": h.casService.GetLogoutURL(),
	})
}

func (h *authHandler) Profile(c *gin.Context) {
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
