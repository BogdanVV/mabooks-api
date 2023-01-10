package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MiddlewareExample(c *gin.Context) {
	fmt.Println("AUTH MIDDLEWARE 1")
	c.Next()
}
