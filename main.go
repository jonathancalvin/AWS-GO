package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonathancalvin/ec2-rds-go/configs"
	"github.com/jonathancalvin/ec2-rds-go/models/database"
)

func init() {
	configs.LoadEnvVariables()
	configs.ConnectToDB()
}
func main() {
	configs.DB.AutoMigrate(&database.Users{})
	router := gin.New()
	router.GET("/user", func(c *gin.Context) {
		var users []database.Users
		result := configs.DB.Find(&users)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Message": result.Error})
			return
		}
		if len(users) == 0 {
			c.JSON(http.StatusOK, gin.H{"Message": "No data found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": users})
	})
	router.Run(":8080")
}
