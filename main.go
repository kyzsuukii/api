package main

import (
	"kyzsuki/api/middlewire"
	"kyzsuki/api/routes/auth"
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

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", auth.Register)
		authGroup.POST("/login", auth.Login)
	}

	r.Run(":8080")
}
