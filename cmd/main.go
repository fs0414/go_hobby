package main

import (
  // "fmt"

  "github.com/fs0414/go_hobby/internal/infrastructure/database"
  // "github.com/fs0414/go_hobby/internal/model"
  "github.com/fs0414/go_hobby/internal/adapter/api/router"
)


func main() {
  database.DbInit()
  router := router.GetRouter()
  router.Run(":8080")
}