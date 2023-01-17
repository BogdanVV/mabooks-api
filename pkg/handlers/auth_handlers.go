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

	_, err = h.services.GetUserByEmail(input.Email)
	if err == nil {
		c.String(http.StatusConflict, "User with this email already exists")
	}

	id, err := h.services.SignUp(input)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	c.String(http.StatusOK, id)
}

func (h *Handlers) login(c *gin.Context) {
	var loginInput models.LoginBody
	err := c.BindJSON(&loginInput)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.GetUserByEmail(loginInput.Email)
	if err != nil {
		c.String(http.StatusNotFound, "The user does not exist")
		return
	}

	err = h.services.VerifyPassword(loginInput.Password, user)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid password")
		return
	}

	response, err := h.services.Login(user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to login")
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

func (h *Handlers) reissueTokens(c *gin.Context) {}
