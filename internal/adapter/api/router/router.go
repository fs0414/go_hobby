package router

import (
	// "github.com/soramar/go_hobby/api/controller"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
  r := gin.Default()
  // r.GET("/", controller.SayHello)
  r.GET("/", func (c *gin.Context)  {
    c.String(200, "hello go hobby")
  })

  return r
}