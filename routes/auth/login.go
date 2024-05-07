package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: implement jwt auth
func Login(c *gin.Context) {
	validator := GetValidator(LoginRequest{})
	req := &LoginRequest{}

	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": validator.DecryptErrors(err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login",
	})
}
