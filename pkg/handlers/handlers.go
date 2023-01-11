package handlers

import (
	"net/http"

	"github.com/bogdanvv/mabooks-api/pkg/middleware"
	"github.com/bogdanvv/mabooks-api/pkg/services"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{services: services}
}

func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/health", checkHealth)

	auth := router.Group("/auth")
	auth.Use(middleware.MiddlewareExample)
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/login", h.login)
		// TODO: REMOVE /token eventually
		auth.POST("/token", h.handleToken)
	}

	api := router.Group("/api")
	{
		api.GET("/health", checkHealth)
	}

	return router
}

func checkHealth(c *gin.Context) {
	c.String(http.StatusOK, "API is healthy")
}
