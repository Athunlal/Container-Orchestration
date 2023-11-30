package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	connectDB()
}

func main() {
	r := gin.Default()

	r.POST("/set", set)
	r.GET("/get", get)

	r.Run(":9000")
}

type User struct {
	gorm.Model // Embedded model with ID, CreatedAt, UpdatedAt, DeletedAt fields
	Name       string
	Email      string
}

func connectDB() {
	psqlInfo := fmt.Sprintf("host=172.17.0.2 port=5432 user=postgres dbname=docker password=pass sslmode=disable")

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("Failed to auto-migrate database")
	}

	DB = db
}

func set(ctx *gin.Context) {
	user := &User{}
	if err := ctx.Bind(user); err != nil {
		ctx.JSON(500, gin.H{
			"Message": "Binding error",
			"Error":   err.Error(),
		})
		return
	}
	fmt.Println("this is the user: ", user.Name)
	result := DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"Message": "creating error",
			"Error":   result.Error,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Succesfully created",
		"Date ":   user,
	})

}

func get(ctx *gin.Context) {

	name := ctx.Query("name")
	user := &User{}

	result := DB.Where("name = ?", name).Find(&user)
	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{
			"Message": "User not found",
			"Error":   result.Error,
		})
		return
	}
	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"Message": "Fetching error",
			"Error":   result.Error,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Succesfully fetching user",
		"Date ":   user,
	})

}
