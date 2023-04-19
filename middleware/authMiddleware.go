package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/padli/go-api-crud/utils"
)

func AuthMiddleware(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")

	if !strings.Contains(bearerToken, "Bearer") {

		c.AbortWithStatusJSON(401, gin.H{
			"msg": "invalid token",
		})

		return
	}

	token := strings.Replace(bearerToken, "Bearer ", "", -1)

	if token == "" {
		c.AbortWithStatusJSON(401, gin.H{
			"msg": "unauthenticated",
		})

		return
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"msg": "unauthenticated",
		})

		return
	}

	c.Set("claimsData", claims)
	c.Next()

}
