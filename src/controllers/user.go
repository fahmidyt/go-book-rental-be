package controllers

import (
	"net/http"

	"github.com/fahmidyt/go-book-rental-be/src/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

var userService = services.InitUserService()

func (ctrl UserController) GetAll(c *gin.Context) {
	data, err := userService.GetAll()

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
