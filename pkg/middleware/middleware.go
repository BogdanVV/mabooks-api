package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authenticate(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	authHeaderSplit := strings.Split(authHeader, " ")
	if len(authHeaderSplit) < 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

	tokenParsed, err := jwt.Parse(authHeaderSplit[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

	if claims, ok := tokenParsed.Claims.(jwt.MapClaims); ok && tokenParsed.Valid {
		c.Set("id", claims["id"].(string))
		c.Set("role", claims["role"].(string))
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

}
