package validations

import (
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/padli/go-api-crud/initializers"
	"github.com/padli/go-api-crud/models"
)

func ExistValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("emailExist", func(fl validator.FieldLevel) bool {
			email := fl.Field().String()

			var user models.User
			initializers.DB.Table("users").Where("email = ?", email).First(&user)

			return user.Email == nil
		})

		v.RegisterValidation("slugUnique", func(fl validator.FieldLevel) bool {
			slug := fl.Field().String()

			var post models.Post
			initializers.DB.Table("posts").Where("slug = ?", slug).First(&post)

			return post.Slug == ""
		})
	}
}

func FileValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("fileExtension", func(fl validator.FieldLevel) bool {
			file, ok := fl.Field().Interface().(*multipart.FileHeader)
			if !ok {
				return false
			}

			extension := filepath.Ext(file.Filename)
			return extension == ".jpg" || extension == ".jpeg" || extension == ".png"
		})

		v.RegisterValidation("fileSize", func(fl validator.FieldLevel) bool {
			file, ok := fl.Field().Interface().(*multipart.FileHeader)
			if !ok {
				return false
			}

			var maxSizeFile int64
			maxSizeFile = 2097152
			return file.Size <= maxSizeFile
		})
	}
}
