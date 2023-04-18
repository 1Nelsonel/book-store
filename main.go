package main

import (
  "github.com/gin-gonic/gin"

  "github.com/1Nelsonel/book-store/models"
  "github.com/1Nelsonel/book-store/controllers" // new
)

func main() {
  r := gin.Default()

  models.ConnectDatabase()

  r.GET("/books", controllers.FindBooks) // new

  r.Run()
}