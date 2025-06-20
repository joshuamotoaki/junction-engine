package application

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.Run()
}