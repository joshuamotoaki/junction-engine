package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tigerappsorg/junction-engine/internal/shared/auth"
	"github.com/tigerappsorg/junction-engine/internal/shared/models"
)

func AuthMiddleware(casService auth.CASService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Try to get token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Redirect(http.StatusFound, casService.GetLoginURL())
			c.Abort()
			return
		}

		// Extract token from "Bearer <token>" format
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":     "Invalid authorization header format",
				"login_url": casService.GetLoginURL(),
			})
			c.Abort()
			return
		}

		// Validate JWT token
		claims, err := casService.ValidateJWT(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":     "Invalid or expired token",
				"login_url": casService.GetLoginURL(),
			})
			c.Abort()
			return
		}

		// Set user information in context
		c.Set("user", claims)
		c.Next()
	}
}

func GetUser(c *gin.Context) (*models.JWTClaims, bool) {
	if user, exists := c.Get("user"); exists {
		if claims, ok := user.(*models.JWTClaims); ok {
			return claims, true
		}
	}
	return nil, false
}
