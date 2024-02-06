package database

import (
  "github.com/joho/godotenv"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "os"
  "fmt"

  // "github.com/fs0414/go_hobby/internal/repository/model"
  "github.com/fs0414/go_hobby/internal/adapter/repository/model"
)

func DbInit() {
  loadEnv()
  fmt.Println("tatatatatttaata")
  
  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
    os.Getenv("POSTGRES_HOST"),
    os.Getenv("POSTGRES_USER"),
    os.Getenv("POSTGRES_PASSWORD"),
    os.Getenv("POSTGRES_DB"),
    os.Getenv("POSTGRES_POST"),
    os.Getenv("POSTGRES_SSLMODE"),
    os.Getenv("POSTGRES_TIMEZONE"),
  )

  database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
      panic("failed to connect to database")
  }

    
  database.AutoMigrate(&model.User{}, &model.Board{})
}

func loadEnv() {
	err := godotenv.Load(".env")
	
	if err != nil {
		fmt.Printf("not loaded env: %v", err)
	}
}