package routes

import (
	"github.com/fahmidyt/go-book-rental-be/src/controllers"
	"github.com/fahmidyt/go-book-rental-be/src/middlewares"
	"github.com/gin-gonic/gin"
)

var bookController = new(controllers.BookController)
var userController = new(controllers.UserController)

func V1Routes(r *gin.RouterGroup) {
	r.GET("/book", bookController.GetAll)

	adminRoute := r.Group("/admin")
	{
		adminRoute.Use(middlewares.AdminAuthMiddleware())

		adminRoute.GET("/user", userController.GetAll)
	}
}
