package controllers

import (
	"fmt"
	"mime/multipart"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/padli/go-api-crud/initializers"
	"github.com/padli/go-api-crud/models"
	"github.com/padli/go-api-crud/validations"
)

type postRequest struct {
	Title    string                `form:"title" binding:"required"`
	Body     string                `form:"body" binding:"required"`
	Image    *multipart.FileHeader `form:"image" binding:"required"`
	ImageUrl string                `form:"image_url"`
}

func PostCreate(c *gin.Context) {
	var req postRequest
	err := c.ShouldBind(&req)
	if err != nil {
		validations.ValidationMsg(err, c)
		return
	}

	// Handle image upload
	file, err := req.Image.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if file == nil {
		c.JSON(404, gin.H{
			"msg": "file required",
		})
		return
	}

	errImg := c.SaveUploadedFile(req.Image, fmt.Sprintf("./public/%s", req.Image.Filename))

	if errImg != nil {
		c.JSON(500, gin.H{
			"msg": "internal server error",
		})
		return
	}

	// Create post data
	post := models.Post{
		Title:    req.Title,
		Body:     req.Body,
		Image:    req.Image.Filename,
		ImageUrl: "http://localhost:5000/public/" + req.Image.Filename,
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"data": post,
		"msg":  "create successfully",
	})
}

func Posts(c *gin.Context) {
	page := c.Query("page")
	if page == "" {
		page = "1"
	}
	perPage := c.Query("perPage")
	if perPage == "" {
		perPage = "10"
	}

	// convert string to int
	pageInt, _ := strconv.Atoi(page)
	perPageInt, _ := strconv.Atoi(perPage)

	if pageInt < 1 {
		pageInt = 1
	}

	// Get posts
	var posts []models.Post
	err := initializers.DB.Table("posts").Offset((pageInt - 1) * perPageInt).Limit(perPageInt).Find(&posts).Error

	if err != nil {
		c.JSON(500, gin.H{
			"msg": "internal server error",
		})
	}

	// Response
	c.JSON(200, gin.H{
		"data":     posts,
		"page":     pageInt,
		"per_page": perPageInt,
	})
}

func Post(c *gin.Context) {
	// Param
	id := c.Param("id")

	// Get posts
	var post models.Post
	// initializers.DB.First(&post, id)
	err := initializers.DB.Table("posts").Where("id = ?", id).Find(&post).Error

	if err != nil {
		c.JSON(500, gin.H{
			"msg": "internal server error",
		})

		return
	}

	// ID in model must *int
	if post.ID == nil {
		c.JSON(404, gin.H{
			"msg": "data not found!",
		})

		return
	}

	// Response
	c.JSON(200, gin.H{
		"data": post,
	})
}

func PostUpdate(c *gin.Context) {
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
		Body:  validation.Body,
	})

	// Response
	c.JSON(200, gin.H{
		"data": post,
	})
}

func PostDelete(c *gin.Context) {
	// Param
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Response
	c.JSON(200, gin.H{
		"msg": "deleted succesfully",
	})
}
