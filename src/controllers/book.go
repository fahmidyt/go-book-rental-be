package controllers

import (
	"net/http"

	"github.com/fahmidyt/go-book-rental-be/src/services"
	"github.com/gin-gonic/gin"
)

type BookController struct {
}

var bookService = services.InitService()

func (ctrl BookController) GetAll(c *gin.Context) {
	data, err := bookService.GetAll()

	if err != nil {
		message := err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully get data!",
		"data":    data,
	})
}

func (ctrl BookController) GetOne(c *gin.Context) {

}
