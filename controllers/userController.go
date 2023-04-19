package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/padli/go-api-crud/initializers"
	"github.com/padli/go-api-crud/models"
	"github.com/padli/go-api-crud/requests"
	"github.com/padli/go-api-crud/validations"
)

func GetAllUser(c *gin.Context) {
	page := c.Query("page")

	if page == "" {
		page = "1"
	}

	perPage := c.Query("perPage")

	if perPage == "" {
		perPage = "5"
	}

	// convert string to int
	perPageInt, _ := strconv.Atoi(perPage)
	pageInt, _ := strconv.Atoi(page)

	if pageInt < 1 {
		pageInt = 1
	}

	var users []models.User

	err := initializers.DB.Table("users").Offset((pageInt - 1) * perPageInt).Limit(perPageInt).Find(&users).Error

	if err != nil {
		c.JSON(500, gin.H{
			"msg": "internal server error",
		})

		return
	}

	c.JSON(200, gin.H{
		"data":     users,
		"page":     pageInt,
		"per_page": perPageInt,
	})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	err := initializers.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if err != nil {
		c.Status(400)
		return
	}

	if user.ID == nil {
		c.JSON(404, gin.H{
			"msg": "data not found!",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}

func CreateUser(c *gin.Context) {
	var req requests.UserRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		validations.ValidationMsg(err, c)
		return
	}

	// userEmailExist := new(models .User)
	// // var userEmailExist models.User
	// initializers.DB.Table("users").Where("email = ?", req.Email).First(&userEmailExist)

	// if userEmailExist.Email != ""  {
	// 	c.JSON(400 , gin.H{
	// 		"msg" : "email already exist",
	// 	})

	// 	return
	// }

	user := models.User{Name: &req.Name, Email: &req.Email, Address: &req.Address}
	result := initializers.DB.Table("users").Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"msg":  "create sucessfully",
		"data": user,
	})
}

func UpdateUser(c *gin.Context) {
	// search data
	id := c.Param("id")
	var user models.User
	initializers.DB.Table("users").Where("id = ?", id).Find(&user)

	if user.ID == nil {
		c.JSON(400, gin.H{
			"msg": "data not found",
		})

		return
	}

	var req struct {
		Name    string `json:"name"`
		Email   string `json:"email" `
		Address string `json:"address"`
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "internal server error",
		})
		return
	}

	if req.Name != "" && req.Name != *user.Name {
		user.Name = &req.Name
		initializers.DB.Model(&user).Update("name", req.Name)
	}
	if req.Email != "" && req.Email != *user.Email {
		user.Email = &req.Email
		initializers.DB.Model(&user).Update("email", req.Email)
	}
	if req.Address != "" && req.Address != *user.Address {
		user.Address = &req.Address
		initializers.DB.Model(&user).Update("address", req.Address)
	}

	// Response
	c.JSON(200, gin.H{
		"msg":  "updated sucessfully",
		"data": user,
	})

}

func DeleteUser(c *gin.Context) {
	// search data
	id := c.Param("id")
	var user models.User
	initializers.DB.Table("users").Where("id = ?", id).Find(&user)

	if user.ID == nil {
		c.JSON(400, gin.H{
			"msg": "data not found",
		})

		return
	}

	// Delete the post
	initializers.DB.Delete(&models.User{}, id)

	// Response
	c.JSON(200, gin.H{
		"msg": "deleted succesfully",
	})
}
