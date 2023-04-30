package router

import (
	"github.com/cleancode/domain"
	"github.com/cleancode/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(gin *gin.Engine, userServie domain.UserService) {

	contrl := handler.NewUserHandler(gin, userServie)

	c := gin.Group("")

	c.POST("/create", contrl.CreateUser)
	c.GET("/:name", contrl.GetUser)

}
