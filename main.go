package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	CustomerName string
	Date string
	Price string
}

func main() {
	r := gin.Default()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	db.AutoMigrate(&Invoice{})
	
  if err != nil {
    panic("failed to connect database")
  }

	r.GET("/", func(ctx *gin.Context) {

		invoice := Invoice{
			CustomerName: "Triple C Collective 2",
			Date: "2024-03-11",
			Price: "1550.00",
		}

		err := db.Create(&invoice)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
			return
		}

		ctx.JSON(http.StatusCreated, invoice)
	})

	r.Run(":8080")
}