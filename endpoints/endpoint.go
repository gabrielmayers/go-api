package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloEndpointHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": 2 + 2,
	})
}

func RedEndpointHandler(c *gin.Context) {
	c.String(200, "red is the new purple")
}

func NewThingsEndpointHandler(c *gin.Context) {
	var json Login

	err := c.ShouldBindJSON(&json)

	if err == nil {
		if json.User == "tony" && json.Password == "1234" {
			c.JSON(200, gin.H{
				"message": "User authenticated!",
			})
		} else {
			c.JSON(400, gin.H{
				"message": "User not found!",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

type Login struct {
	User     string
	Password string
}
