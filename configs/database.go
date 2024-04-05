package configs

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error = nil;
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(dsn)
	if err != nil {
		fmt.Println("Failed to connect to database")
		log.Fatal("Fail")
	}
}