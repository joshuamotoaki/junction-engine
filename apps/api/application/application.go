package application

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tigerappsorg/junction-engine/config"
	"github.com/tigerappsorg/junction-engine/database"
)

func Run(cfg *config.Config) {
	db, err := database.NewNeo4j(cfg.NEO4J_URI, cfg.NEO4J_USERNAME, cfg.NEO4J_PASSWORD)
	if err != nil {
		log.Fatal("Failed to connect to Neo4j:", err)
	}
	defer db.Close(context.Background())

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		ctx := context.Background()
		if err := db.HealthCheck(ctx); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "unhealthy",
				"error":  "database connection failed",
			})
			return
		}

		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.Run()
}
