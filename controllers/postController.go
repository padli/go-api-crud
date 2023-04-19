package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/padli/go-api-crud/initializers"
	"github.com/padli/go-api-crud/models"
	"github.com/padli/go-api-crud/requests"
	"github.com/padli/go-api-crud/validations"
)

func PostCreate(c *gin.Context) {
	var req requests.PostRequest
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

	// Check file extension
	extension := filepath.Ext(req.Image.Filename)
	if extension != ".jpg" && extension != ".jpeg" && extension != ".png" {
		c.JSON(400, gin.H{
			"msg": "invalid file type",
		})
		return
	}

	// Check file size
	var maxSizeFile int64
	maxSizeFile = 2097152
	if req.Image.Size > maxSizeFile {
		c.JSON(400, gin.H{
			"msg": "file size too large",
		})
		return
	}

	newFileName := fmt.Sprintf("%s_%s%s", time.Now().Format("20060102150405"), strings.Split(req.Image.Filename, ".")[0], filepath.Ext(req.Image.Filename))
	errImg := c.SaveUploadedFile(req.Image, fmt.Sprintf("./public/%s", newFileName))

	if errImg != nil {
		c.JSON(500, gin.H{
			"msg": "internal server error",
		})
		return
	}

	// url := c.Request.URL
	protocol := "http"
	if c.Request.TLS != nil {
		protocol = "https"
	}
	host := c.Request.Host
	newURL := fmt.Sprintf("%s://%s/public/%s", protocol, host, newFileName)

	// Create post data
	post := models.Post{
		Title:    req.Title,
		Slug:     req.Slug,
		Body:     req.Body,
		Image:    newFileName,
		ImageUrl: newURL,
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
			"msg": err.Error(),
		})
		return
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
	slug := c.Param("slug")

	// Get posts
	var post models.Post
	// initializers.DB.First(&post, id)
	err := initializers.DB.Table("posts").Where("slug = ?", slug).Find(&post).Error

	if err != nil {
		c.JSON(500, gin.H{
			"msg": "internal server error",
		})

		return
	}

	// ID in model must *int
	if post.Slug == "" {
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
	slug := c.Param("slug")

	// Get data req body
	var req requests.PostRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		validations.ValidationMsg(err, c)
		return
	}

	// Find the  post were updating
	var post models.Post
	initializers.DB.First(&post, slug)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: req.Title,
		Body:  req.Body,
		Slug:  req.Slug,
		Image: req.Image.Filename,
	})

	// Response
	c.JSON(200, gin.H{
		"data": post,
	})
}

func PostDelete(c *gin.Context) {
	// Param
	slug := c.Param("slug")

	// Get the post
	var post models.Post
	if err := initializers.DB.First(&post, slug).Error; err != nil {
		c.JSON(404, gin.H{"msg": "post not found"})
		return
	}

	// Delete the post
	if err := initializers.DB.Delete(&post).Error; err != nil {
		c.JSON(400, gin.H{"msg": "failed to delete post"})
		return
	}

	// Delete the file
	filePath := fmt.Sprintf("./public/%s", post.Image)
	if err := os.Remove(filePath); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	// Response
	c.JSON(200, gin.H{"msg": "deleted successfully"})
}
