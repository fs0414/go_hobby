package router
// "github.com/fs0414/go_hobby/internal/adapter/api/router"

import (
	"github.com/gin-gonic/gin"

  "github.com/fs0414/go_hobby/internal/usecase"
  "github.com/fs0414/go_hobby/internal/adapter/repository"
)

func GetRouter() *gin.Engine {
  r := gin.Default()
  r.GET("/", func (c *gin.Context)  {
    c.String(200, "hello go hobby")
  })

  userRepo := repository.UserRepositoryFactory()
  userUseCase := usecase.UserUseCaseFactory(userRepo)

  boardRepo := repository.BoardRepositoryFactory()
  boardUseCase := usecase.BoardUseCaseFactory(boardRepo)

  apiRouter := r.Group("/api")
  {
    apiRouter.GET("/users", userUseCase.FetchUsers)
    apiRouter.POST("/users", userUseCase.CreateUser)
    
    apiRouter.GET("/boards", boardUseCase.FetchBoards)
    apiRouter.POST("/boards", boardUseCase.CreateBoard)
  }

  return r
}