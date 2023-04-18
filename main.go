package main

import (
  "github.com/gin-gonic/gin"

  "github.com/rahmanfadhil/gin-bookstore/models" // new
)

func main() {
  r := gin.Default()

  models.ConnectDatabase() // new

  r.Run()
}