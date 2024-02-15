package usecase
// "github.com/fs0414/go_hobby/internal/usecase"

import (
	// "fmt"

	"gorm.io/gorm"
	"github.com/gin-gonic/gin"

	"github.com/fs0414/go_hobby/internal/adapter/repository"
)

var db *gorm.DB

type User struct {}

type UserUseCase struct {
	repo repository.UserRepository
}

func UserUseCaseFactory(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase)FetchUsers(c *gin.Context) {
	
	users, err := uc.repo.GetUsers()
	if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, users)
}