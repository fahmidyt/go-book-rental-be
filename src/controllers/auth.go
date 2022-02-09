package controllers

import (
	"net/http"

	"github.com/fahmidyt/go-book-rental-be/src/models"
	"github.com/fahmidyt/go-book-rental-be/src/types"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

var userModel = new(models.User)

func (ctl AuthController) Login(c *gin.Context) {
	var body types.LoginForm

	// if validation error
	if validationErr := c.ShouldBindJSON(&body); validationErr != nil {
		message := "Failed to login. Please check again your data."
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

}
