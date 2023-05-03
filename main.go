package main

import (
	"github.com/gin-gonic/gin"
	"github.com/padli/go-api-crud/controllers"
	"github.com/padli/go-api-crud/initializers"
	"github.com/padli/go-api-crud/middleware"
	"github.com/padli/go-api-crud/validations"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	// Custom Validations
	validations.ExistValidation()
	validations.FileValidation()

	r := gin.Default()
	r.Static("/public", "./public")

	auth := r.Group("", middleware.AuthMiddleware)

	// AUTH ENDPOINT
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequiredAuth, controllers.Validate)

	// POST ENDPOINT
	auth.GET("/posts", controllers.Posts)
	r.GET("/posts/:slug", controllers.Post)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.POST("/posts", controllers.PostCreate)

	// CATEGORY ENPOINT
	r.POST("/category", controllers.CategoryCreate)
	r.GET("/category", controllers.Categories)
	r.GET("/category/:id", controllers.Category)
	r.PUT("/category/:id", controllers.CategoryUpdate)
	r.DELETE("/category/:id", controllers.CategoryDelete)

	// RUN APP
	r.Run()
}
