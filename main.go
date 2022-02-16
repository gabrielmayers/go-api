package main

import (
	"github.com/gin-gonic/gin"

	"go-api/endpoints"
)

func main() {
	router := gin.Default()

	router.GET("/hello", endpoints.HelloEndpointHandler)
	router.GET("/red", endpoints.RedEndpointHandler)
	router.POST("/new-things", endpoints.NewThingsEndpointHandler)

	router.Run()
}
