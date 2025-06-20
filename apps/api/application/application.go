package application

import (
	"github.com/gin-gonic/gin"
	"github.com/tigerappsorg/junction-engine/config"
)

func Run(cfg *config.Config) {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.Run()
}