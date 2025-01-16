package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RunMigrations(client *mongo.Client, dbName string) error  {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	db := client.Database(dbName)


	_, err := db.Collection("users").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "email", Value: 1}, {Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	})

	if err != nil {
		log.Printf("Error creating unique index on email: %v", err)
		return err
	}

	log.Println("Migraions completed successfully")
	return nil
}