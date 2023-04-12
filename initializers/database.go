package initializers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB *gorm.DB
func ConnectDB(){
	// var err error
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = db
	
	

	if err != nil {
		log.Println("Failed connect to database")
	} else {
		log.Println("connected to database")
	}
}