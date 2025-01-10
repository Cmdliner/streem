package repository

import (
	"context"
	"time"

	"github.com/Cmdliner/streem/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OtpRepository struct {
	collection *mongo.Collection
}

func NewOtpRepository(db *mongo.Database) *OtpRepository{
	return &OtpRepository{
		collection: db.Collection("otps"),
	}
}

func (r *OtpRepository) Create(user *model.User, kind string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	newOtp := model.OTP {
		User: user.ID,
		CreatedAt: time.Now(),
		Kind: kind,
		Code: "1234", // !todo => change this to an actual rand value
	}

	otp, err := r.collection.InsertOne(ctx, newOtp)
	newOtp.ID = otp.InsertedID.(primitive.ObjectID)
	if err != nil {
		return "", err
	}

	return newOtp.Code, nil
}

func (r *OtpRepository) GetOne(user *model.User, code, kind string) (*model.OTP, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	var otp model.OTP

	err := r.collection.FindOne(ctx, bson.M{"user": user.ID, "code": code, "kind": kind}).Decode(&otp)
	if err != nil {
		return nil, err
	}

	return &otp, nil
}

func (r *OtpRepository) Delete() {}

func (r *OtpRepository) DeleteMany() {}