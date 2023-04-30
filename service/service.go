package service

import (
	"context"
	"time"

	"github.com/cleancode/domain"
)

type userService struct {
	userRepository domain.UserRepository
	timeout        time.Duration
}

func NewUserService(userRepository domain.UserRepository, timeout time.Duration) domain.UserService {
	return &userService{
		userRepository: userRepository,
		timeout:        timeout,
	}
}

func (u *userService) CreateUser(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, u.timeout*time.Second)
	defer cancel()
	return u.userRepository.CreateUser(ctx, user)
}

func (u *userService) GetUser(c context.Context, name string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, u.timeout*time.Second)
	defer cancel()

	return u.userRepository.GetUser(ctx, name)
}
