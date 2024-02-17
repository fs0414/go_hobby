package repository_interface

import (
	"github.com/fs0414/go_hobby/internal/infrastructure/model"
)

type UserRepositoryIf interface {
	GetUsers() ([]model.User, error)
	CreateUser(user model.User) (model.User, error)
}