package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/padli/go-api-crud/initializers"
	"github.com/padli/go-api-crud/models"
)

// create post

func PostCreate (c *gin.Context) {
	// Get data req body
	var body struct{
		Body string
		Title string
	}
	c.Bind(&body)

	// Creat post data
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"data": post,
	})
}

func Posts(c *gin.Context){
	// Get posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Response
	c.JSON(200, gin.H{
		"data" : posts,
	})
}

func Post(c *gin.Context){
	// Param
	id := c.Param("id")

	// Get posts
	var post models.Post
	initializers.DB.First(&post, id)

	// Response
	c.JSON(200, gin.H{
		"data" : post,
	})
}


func PostUpdate(c *gin.Context){
	// Param
	id := c.Param("id")


	// Get data req body
	var body struct{
		Body string
		Title string
	}
	c.Bind(&body)

	// Find the  post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	// Response
	c.JSON(200, gin.H{
		"data" : post,
	})
}

func PostDelete(c *gin.Context){
	// Param
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Response
	c.JSON(200, gin.H{
		"msg" : "deleted succesfully",
	})
}