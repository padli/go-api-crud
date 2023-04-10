package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/padli/go-api-crud/initializers"
	"github.com/padli/go-api-crud/models"
)

type postValidation struct{
	Title string  `json:"title" binding:"required"`
	Body string   `json:"body" binding:"required"`
}



// create post
func PostCreate (c *gin.Context) {

	var validation postValidation
	err := c.ShouldBindJSON(&validation)
	if err != nil {
		
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("%s : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors" : errorMessages,
		})
		return
	}
	// Creat post data
	post := models.Post{Title: validation.Title, Body: validation.Body}
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
	var validation postValidation
	err := c.ShouldBindJSON(&validation)
	if err != nil {
		
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("%s : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors" : errorMessages,
		})
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