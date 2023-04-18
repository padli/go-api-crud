package main

import (
	"github.com/padli/go-api-crud/initializers"
	"github.com/padli/go-api-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
