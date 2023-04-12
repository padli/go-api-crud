package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/padli/go-api-crud/initializers"
	"github.com/padli/go-api-crud/models"
	"github.com/padli/go-api-crud/validations"
)

type postRequest struct{
	Title string  `json:"title" form:"title" binding:"required"`
	Body string   `json:"body" form:"body" binding:"required"`
}



// create post
func PostCreate (c *gin.Context) {

	var req postRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		validations.ValidationMsg(err, c)
		return
	}
	
	// Creat post data
	post := models.Post{Title: req.Title, Body: req.Body}
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
	err := initializers.DB.Table("posts").Find(&posts).Error

	if err != nil {
		c.JSON(500, gin.H{
			"msg" : "internal server error",
		})
	}

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
	// initializers.DB.First(&post, id)
	err := initializers.DB.Table("posts").Where("id = ?", id).Find(&post).Error

	if err != nil  {
		c.JSON(500, gin.H{
			"msg" : "internal server error",
		})

		return
	}

	// ID in model must *int
	if  post.ID == nil {
		c.JSON(404, gin.H{
			"msg" : "data not found!",
		})

		return
	}

	// Response
	c.JSON(200, gin.H{
		"data" : post,
	})
}


func PostUpdate(c *gin.Context){
	// Param
	id := c.Param("id")


	// Get data req body
	var validation postRequest
	err := c.ShouldBindJSON(&validation)
	if err != nil {
		validations.ValidationMsg(err, c)
		return
	}

	
	// Find the  post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: validation.Title,
		Body: validation.Body,
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