package repository
// "github.com/fs0414/go_hobby/internal/adapter/repository"

import (
	"github.com/fs0414/go_hobby/internal/adapter/repository/interface"
	"github.com/fs0414/go_hobby/internal/infrastructure/database"
	"github.com/fs0414/go_hobby/internal/infrastructure/model"
)

type UserRepository struct {}

func UserRepositoryFactory() repository_interface.UserRepositoryIf {
	return &UserRepository{}
}

func (repo *UserRepository) GetUsers() ([]model.User, error) {
	db := database.GetDb()
	var users []model.User
    result := db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (repo *UserRepository) CreateUser(user model.User) (model.User, error) {
	db := database.GetDb()
	result := db.Create(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}