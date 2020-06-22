package main

import (
	controller "github.com/bookstore/controllers"
	"github.com/bookstore/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectToDB()

	r.GET("/books", controller.FindBooks)
	r.POST("/books", controller.CreateBook)

	r.Run()
}
