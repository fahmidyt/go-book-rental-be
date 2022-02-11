package routes

import (
	"github.com/fahmidyt/go-book-rental-be/src/controllers"
	"github.com/gin-gonic/gin"
)

var bookController = new(controllers.BookController)

func V1Routes(r *gin.RouterGroup) {
	r.GET("/book", bookController.GetAll)
}
