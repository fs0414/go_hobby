package usecase
// "github.com/fs0414/go_hobby/internal/usecase"

import (
	// "fmt"

	"github.com/gin-gonic/gin"

	"github.com/fs0414/go_hobby/internal/adapter/repository/interface"
	"github.com/fs0414/go_hobby/internal/infrastructure/model"
)

type UserUseCase struct {
	repo repository_interface.UserRepositoryIf
}

func UserUseCaseFactory(repo repository_interface.UserRepositoryIf) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) FetchUsers(c *gin.Context) {
	users, err := uc.repo.GetUsers()
	if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, users)
}

func (uc *UserUseCase) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

	newUser, err := uc.repo.CreateUser(user)

	if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

	c.JSON(201, newUser)
}