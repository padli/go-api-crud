package validations

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/padli/go-api-crud/initializers"
	"github.com/padli/go-api-crud/models"
)

func EmailExistValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("emailExist", func(fl validator.FieldLevel) bool {
			email := fl.Field().String()

			var user models.User
			initializers.DB.Table("users").Where("email = ?", email).First(&user)

			return user.Email == nil
		})
	}
}

func SlugExistValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("slugUnique", func(fl validator.FieldLevel) bool {
			slug := fl.Field().String()

			var post models.Post
			initializers.DB.Table("posts").Where("slug = ?", slug).First(&post)

			return post.Slug == ""
		})
	}
}
