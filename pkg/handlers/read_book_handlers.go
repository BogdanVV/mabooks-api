package handlers

import (
	"fmt"
	"net/http"

	models "github.com/bogdanvv/mabooks-api"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) CreateBook(c *gin.Context) {
	var readBookInput models.ReadBookInput
	userId := c.MustGet("id").(string)

	err := c.Bind(&readBookInput)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	readBook, err := h.services.CreateReadBook(userId, readBookInput)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, readBook)
}

func (h *Handlers) GetBookById(c *gin.Context) {
	bookId := c.Param("bookId")
	userId := c.MustGet("id").(string)

	book, err := h.services.GetBookById(bookId)
	if err != nil {
		c.String(http.StatusNotFound, "Book with such id does not exist")
		fmt.Println(err.Error())
		return
	}

	if book.UserId == userId {
		c.JSON(http.StatusOK, book)
		return
	} else {
		c.String(http.StatusUnauthorized, "The book does not belong to the user")
		return
	}
}

func (h *Handlers) GetAllBooksByUserId(c *gin.Context) {
	id := c.MustGet("id").(string)

	books, err := h.services.GetAllBooksByUserId(id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not find the books")
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *Handlers) UpdateBook(c *gin.Context) {
	bookId := c.Param("bookId")
	var book models.ReadBookInput
	c.Bind(&book)

	bookResult, err := h.services.UpdateBook(bookId, book)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, bookResult)
}

func (h *Handlers) DeleteBook(c *gin.Context) {
	bookId := c.Param("bookId")
	id, err := h.services.DeleteBook(bookId)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.String(http.StatusOK, id)
}
