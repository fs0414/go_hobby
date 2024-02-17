package router

import (
    "github.com/gin-gonic/gin"

    "github.com/fs0414/go_hobby/internal/usecase"
	"github.com/fs0414/go_hobby/internal/adapter/repository/user"
)

func SetUserRoutes(rg *gin.RouterGroup) {
	uc := usecase.UserUseCaseFactory(repository.UserRepositoryFactory())

	rg.GET("/users", uc.FetchUsers)
	rg.POST("/users", uc.CreateUser)
}