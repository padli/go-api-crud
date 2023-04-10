package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/padli/go-api-crud/initializers"
	"github.com/padli/go-api-crud/models"
)

func CategoryCreate(c *gin.Context){
	var body struct{
		Title string 
		Desc string
	}
	c.Bind(&body)

	category := models.Category{Title: body.Title, Desc: body.Desc}
	result := initializers.DB.Create(&category)

	

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"msg" : "create succesfully",
		"data" : category,
	})
	
}

func Categories(c *gin.Context){
	var categories []models.Category
	initializers.DB.Find(&categories)

	c.JSON(200, gin.H{
		"data" : categories,
	})
}

func Category (c *gin.Context){
	id := c.Param("id")
	
	var category models.Category
	initializers.DB.First(&category, id)

	c.JSON(200, gin.H{
		"data" : category,
	})

}

func CategoryUpdate(c *gin.Context){
	id := c.Param("id")

	var body struct{
		Title string 
		Desc string
	}
	c.Bind(&body)

	var category models.Category
	initializers.DB.First(&category, id)

	initializers.DB.Model(&category).Updates(models.Category{
		Title: body.Title,
		Desc: body.Desc,
	})

	c.JSON(200, gin.H{
		"updated" : category,
	})

}

func CategoryDelete(c *gin.Context){
	id := c.Param("id")

	initializers.DB.Delete(&models.Category{}, id)

	c.JSON(200, gin.H{
		"msg" : "deleted successfuly",
	})
}