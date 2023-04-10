package main

import (
	"github.com/gin-gonic/gin"
	"github.com/padli/go-api-crud/controllers"
	"github.com/padli/go-api-crud/initializers"
)


func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}


func main(){
	r := gin.Default()

	// POST ENDPOINT
	r.GET("/posts", controllers.Posts)
	r.GET("/posts/:id", controllers.Post)
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