package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type OTP struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User primitive.ObjectID `json:"user" bson:"user"`
	Code string `json:"code" bson:"code"`
	Kind string `bson:"kind" json:"kind"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}