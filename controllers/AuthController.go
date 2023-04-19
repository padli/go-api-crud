package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/padli/go-api-crud/initializers"
	"github.com/padli/go-api-crud/models"
	"github.com/padli/go-api-crud/requests"
	"github.com/padli/go-api-crud/utils"
)

func Login(c *gin.Context) {
	var req requests.LoginRequest

	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"msg": err.Error(),
		})
	}

	var user models.User
	err := initializers.DB.Table("users").Where("email = ?", req.Email).First(&user).Error

	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"msg": "credential error",
		})

		return
	}

	// if user.Email != nil {
	// 	c.AbortWithStatusJSON(404, gin.H{
	// 		"msg": "email invalid",
	// 	})

	// 	return
	// }

	// check password
	if req.Password != "12345" {
		c.AbortWithStatusJSON(404, gin.H{
			"msg": "wrong password",
		})

		return
	}

	claims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := utils.GenerateToken(&claims)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"msg": "failed generate token",
		})

		return
	}

	c.JSON(200, gin.H{
		"msg":   "login successfully",
		"token": token,
	})
}
