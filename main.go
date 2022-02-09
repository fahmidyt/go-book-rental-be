package main

import (
	"log"
	"os"

	"github.com/fahmidyt/go-book-rental-be/src/db"
	"github.com/fahmidyt/go-book-rental-be/src/middlewares"
	"github.com/fahmidyt/go-book-rental-be/src/models"
	"github.com/fahmidyt/go-book-rental-be/src/routes"

	"gorm.io/gorm"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// TODO: make it simple
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Book{}, &models.Role{}, &models.User{}, &models.RentedBook{}, &models.RentedBookDetail{}, &models.UserDetail{})
}

const HTML_GLOB = "./public/html/*"

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error: failed to load file env file")
	}

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	// run default server of gin
	r := gin.Default()

	// main middlewares here
	r.Use(middlewares.CORSMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	// every endpoint that start with /v1 will use auth middleware
	v1Group := r.Group("/v1", middlewares.AuthMiddleware())

	db := db.OpenConnection()
	AutoMigrate(db)

	// all of the routes declared here
	routes.PublicRoutes(r)
	routes.V1Routes(v1Group)

	r.LoadHTMLGlob(HTML_GLOB)
	r.Static("/public", "./public")

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	port := os.Getenv("PORT")
	log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))

	r.Run(":" + port)

}
