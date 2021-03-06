package routes

import (
	"net/http"

	"github.com/fahmidyt/go-book-rental-be/src/controllers"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Book Rental API Service",
			"author":  "fahmidyt",
			"source":  "https://github.com/fahmidyt/go-book-rental-be",
		})
	})

	// public auth routes
	auth := r.Group("/auth")
	{
		// declare auth controller
		authController := new(controllers.AuthController)

		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
	}
}
