package repository
// "github.com/fs0414/go_hobby/internal/adapter/repository"

import (
	"fmt"

	"github.com/fs0414/go_hobby/internal/infrastructure/database"
	"github.com/fs0414/go_hobby/internal/infrastructure/model"
)

type UserRepository interface {
	GetUsers() ([]model.User, error)
}

type UserRepositoryImpl struct {
	UserRepository
}

func UserRepositoryFactory() UserRepository {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl)GetUsers() ([]model.User, error) {
	var users []model.User
	db := database.Db
    result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println("testtest")
	return users, nil
}