package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/health", checkHealth)
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/login")
	}

	router.Run()
	fmt.Printf("Server is running on port :%s", "8080")
}

func checkHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API is healthy"})
}
