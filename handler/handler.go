package handler

import (
	"context"
	"net/http"

	"github.com/cleancode/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	gin         *gin.Engine
	userService domain.UserService
}

func NewUserHandler(gin *gin.Engine, userService domain.UserService) *UserHandler {
	return &UserHandler{
		gin:         gin,
		userService: userService,
	}
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var user *domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = primitive.NewObjectID()

	if err := u.userService.CreateUser(context.Background(), user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "inserted")
}

func (u *UserHandler) GetUser(c *gin.Context) {
	name := c.Param("name")

	user, err := u.userService.GetUser(c, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)

}
