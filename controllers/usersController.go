package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/padli/go-api-crud/initializers"
	"github.com/padli/go-api-crud/models"
	"github.com/padli/go-api-crud/validations"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	// req body
	var req struct {
		Name            string `json:"name" binding:"required"`
		Email           string `json:"email" binding:"required,email"`
		Password        string `json:"password" binding:"required,min=5,alphanum"`
		ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	}
	err := c.ShouldBind(&req)
	if err != nil {
		validations.ValidationMsg(err, c)
		return
	}

	// hash password
	hash, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	hashString := string(hash)

	if errHash != nil {
		c.JSON(500, gin.H{
			"msg": "failed hash password",
		})

		return
	}

	// create user
	user := models.User{Name: &req.Name, Email: &req.Email, Password: &hashString}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"msg": "failed create user",
		})

		return
	}

	// response
	c.JSON(200, gin.H{
		"data": user,
		"msg":  "created user",
	})

}

func Login(c *gin.Context) {
	// req body
	// var req requests.UserRequest
	var req struct {
		Email    string
		Password string
	}
	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(500, gin.H{
			"msg": err.Error(),
		})

		return
	}

	// check user
	var user models.User
	initializers.DB.First(&user, "email = ?", req.Email)

	if user.ID == nil {
		c.JSON(404, gin.H{
			"msg": "user not found",
		})
		return
	}

	// compare password
	errPassword := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(req.Password))
	if errPassword != nil {
		c.JSON(400, gin.H{
			"msg": "wrong password!",
		})
		return
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, errToken := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if errToken != nil {
		c.JSON(400, gin.H{
			"msg": "invalid create token",
		})
		return
	}

	// cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	// result
	c.JSON(200, gin.H{
		"name":  user.Name,
		"token": tokenString,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	name := user.(models.User).Name
	c.JSON(200, gin.H{
		"hello": name,
	})
}
