package main

import (
	"kyzsuki/api/middlewire"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.SetTrustedProxies(nil)
	r.Use(middlewire.DefaultStructuredLogger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}
