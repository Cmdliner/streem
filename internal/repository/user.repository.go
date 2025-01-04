package repository

import (
	"context"
	"time"

	"github.com/Cmdliner/streem/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) (*UserRepository) {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	user, err := r.collection.InsertOne(ctx, user)
	if err != nil  {
		return nil, err
	}

	user.ID
}

func (r * UserRepository) Update() {}

func (r * UserRepository) Delete() {}