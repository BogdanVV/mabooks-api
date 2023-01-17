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
	router.Use(middleware.CORSMiddleware())

	router.GET("/health", checkHealth)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/login", h.login)
		// auth.POST("/tokens", h.reissueTokens)
		auth.POST("/token", h.handleToken)
	}

	api := router.Group("/api")
	api.Use(middleware.Authenticate)
	{
		readBooks := api.Group("/read-books")
		{
			readBooks.POST("/", h.CreateBook)
			readBooks.GET("/", h.GetAllBooksByUserId)
			readBooks.GET("/:bookId", h.GetBookById)
			readBooks.PATCH("/:bookId", h.UpdateBook)
			readBooks.DELETE("/:bookId", h.DeleteBook)
		}
	}

	return router
}

func checkHealth(c *gin.Context) {
	c.String(http.StatusOK, "API is healthy")
}
