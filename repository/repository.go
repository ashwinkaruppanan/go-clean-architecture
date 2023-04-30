package repository

import (
	"context"

	"github.com/cleancode/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	database   *mongo.Database
	collection string
}

func NewUserRepository(db *mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}

}

func (u *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	_, err := u.database.Collection(u.collection).InsertOne(ctx, user)
	return err
}

func (u *userRepository) GetUser(ctx context.Context, name string) (domain.User, error) {

	var usr domain.User
	err := u.database.Collection(u.collection).FindOne(ctx, bson.M{"name": name}).Decode(&usr)
	return usr, err
}
