package main

import (
  "github.com/soramar/go_hobby/model"
  "github.com/soramar/go_hobby/router"
)

func main() {
  model.DbInit()

  router := router.GetRouter()
  router.Run(":8080")
}