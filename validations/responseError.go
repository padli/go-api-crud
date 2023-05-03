package validations

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type errorMsg struct {
	Field   string `json:"param"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "email":
		return "Invalid email " + fe.Param()
	case "emailExist":
		return "Email already exist " + fe.Param()
	case "slugUnique":
		return "Slug already exist " + fe.Param()
	case "fileExtension":
		return "Invalid file type " + fe.Param()
	case "fileSize":
		return "file size too large " + fe.Param()
	case "eqfield":
		return "Password and Confirm Password do not match" + fe.Param()
	}
	return "Unknown error"
}

func ValidationMsg(err error, c *gin.Context) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]errorMsg, len(ve))
		for i, fe := range ve {
			out[i] = errorMsg{fe.Field(), getErrorMsg(fe)}
		}
		c.JSON(400, gin.H{"errors": out})
	}
}
