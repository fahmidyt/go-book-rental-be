package helpers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/fahmidyt/go-book-rental-be/src/services"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

func VerifyAccesToken(c *gin.Context) {
	token, ok := getToken(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	if !jwtToken.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		return
	}

	c.Next()
}

func getToken(c *gin.Context) (currentToken string, ok bool) {
	if cookie, err := c.Request.Cookie("token"); err == nil {
		currentToken = cookie.Value
	} else {
		getHeader := c.Request.Header.Get("Authorization")

		if len(getHeader) <= 0 {
			return currentToken, false
		}

		splitToken := strings.Split(getHeader, " ")

		if len(splitToken) == 2 && (splitToken[0] == "Bearer" || splitToken[0] == "JWT") {
			return splitToken[1], true
		}
	}

	return currentToken, false
}

func extractMetadata(token *jwt.Token) (*services.TokenDetails, error) {
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)

		if !ok {
			err := errors.New("cannot find access_uuid from token")
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}

		return &services.TokenDetails{
			AccessUUID: accessUUID,
			UserID:     uint(userId),
		}, nil
	}
	err := errors.New("token invalid")
	return nil, err
}
