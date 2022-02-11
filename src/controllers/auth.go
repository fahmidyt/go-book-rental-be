package controllers

import (
	"net/http"

	"github.com/fahmidyt/go-book-rental-be/src/services"
	"github.com/fahmidyt/go-book-rental-be/src/types"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

// declare auth controller
var authService = new(services.AuthService)

func (ctrl AuthController) Login(c *gin.Context) {
	var body types.LoginForm

	// if validation error
	if validationErr := c.ShouldBindJSON(&body); validationErr != nil {
		message := "Failed to login. Please check again your data."
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	user, token, err := authService.Login(body)

	if err != nil {
		message := err.Error()
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully login!",
		"user":    user,
		"token":   token,
	})
}

func (ctrl AuthController) Register(c *gin.Context) {
	var body types.RegisterForm

	// if validation error
	if validationErr := c.ShouldBindJSON(&body); validationErr != nil {
		message := "Your data is invalid. Please check your information again"
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	user, err := authService.Register(body)

	if err != nil {
		message := err.Error()
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully register",
		"data":    user,
	})
}
