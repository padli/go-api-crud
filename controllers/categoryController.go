package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/padli/go-api-crud/initializers"
	"github.com/padli/go-api-crud/models"
)

type categoryValidation struct{
	Title string  `json:"title" binding:"required"`
	Desc string   `json:"desc" binding:"required"`
}

func CategoryCreate(c *gin.Context){

	var validation categoryValidation
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

	category := models.Category{Title: validation.Title, Desc: validation.Desc}
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

	var validation categoryValidation
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

	var category models.Category
	initializers.DB.First(&category, id)

	initializers.DB.Model(&category).Updates(models.Category{
		Title: validation.Title,
		Desc: validation.Desc,
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