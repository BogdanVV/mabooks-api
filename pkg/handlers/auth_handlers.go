package handlers

import (
	"net/http"
	"strings"

	models "github.com/bogdanvv/mabooks-api"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) signUp(c *gin.Context) {
	var input models.SignUpBody
	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.SignUp(input)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	c.String(http.StatusOK, id)
}

func (h *Handlers) login(c *gin.Context) {
	var loginInput models.LoginBody
	err := c.Bind(&loginInput)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.services.Login(loginInput)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handlers) handleToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	authHeaderSplit := strings.Split(authHeader, " ")
	if len(authHeaderSplit) < 2 {
		c.String(http.StatusInternalServerError, "Invalid 'Authorization' header")
		return
	}

	h.services.HandleToken(authHeaderSplit[1])
}
