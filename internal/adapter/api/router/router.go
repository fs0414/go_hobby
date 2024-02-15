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

  apiRouter := r.Group("/api")
  {
    apiRouter.GET("/users", userUseCase.FetchUsers)
    // apiRouter.GET("/users",usecase.FetchUsers)
  }

  return r
}