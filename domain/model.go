package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, email string) (User, error)
}

type UserService interface {
	CreateUser(c context.Context, user *User) error
	GetUser(c context.Context, email string) (User, error)
}
