package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tigerappsorg/junction-engine/database"
)

type healthHandler struct {
	db database.Neo4jDB
}

type HealthHandler interface{
	Check(c *gin.Context)
	DatabaseStatus(c *gin.Context)
}

func NewHealthHandler(db database.Neo4jDB) HealthHandler {
	return &healthHandler{
		db: db,
	}
}

func (h *healthHandler) Check(c *gin.Context) {
	ctx := context.Background()
	if err := h.db.HealthCheck(ctx); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unhealthy",
			"error":  "database connection failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"database": "connected",
	})
}

func (h *healthHandler) DatabaseStatus(c *gin.Context) {
	ctx := context.Background()
	if err := h.db.HealthCheck(ctx); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"database": "disconnected",
			"error":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"database": "connected",
		"status":   "healthy",
	})
}
