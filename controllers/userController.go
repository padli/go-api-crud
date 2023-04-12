package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/padli/go-api-crud/initializers"
	"github.com/padli/go-api-crud/models"
	"github.com/padli/go-api-crud/validations"
)

type userRequest struct{
	Name 	string 	`json:"name" binding:"required"`
	Email 	string 	`json:"email" binding:"required,email"`
	Address string 	`json:"address" binding:"required"`
}



func GetAllUser (c *gin.Context){
	var users []models.User

	err := initializers.DB.Table("users").Find(&users).Error

	if err != nil {
		c.JSON(500, gin.H{
			"msg" : "internal server error",
		})

		return
	}

	c.JSON(200, gin.H{
		"data" : users,
	})
}

func CreateUser(c *gin.Context){
	var req userRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		validations.ValidationMsg(err, c)
		return
	}

	user := models.User{Name: req.Name, Email: req.Email, Address: req.Address}
	result := initializers.DB.Table("users").Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"msg" : "create sucessfully",
		"data" : user,
	})
}