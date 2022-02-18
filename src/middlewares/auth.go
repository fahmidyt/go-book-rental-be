package middlewares

import (
	"net/http"

	"github.com/fahmidyt/go-book-rental-be/src/db"
	"github.com/fahmidyt/go-book-rental-be/src/helpers"
	"github.com/fahmidyt/go-book-rental-be/src/models"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return helpers.VerifyAccesToken
}

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("User").(map[string]interface{})

		result := new(models.Role)
		res := db.GetDB().Model(&result).First(result, models.Role{Name: "Admin"})

		if res.Error != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"message": "Role Admin cannot be found in database!",
			})
			return
		}

		if uint(user["RoleId"].(float64)) != result.ID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		c.Next()
	}
}
