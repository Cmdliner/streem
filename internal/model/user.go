package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Email string `json:"email" bson:"email"`
	Password string `json:"-" bson:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

}