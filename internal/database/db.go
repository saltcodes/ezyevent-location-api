package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Config() *mongo.Database {
	ctx := context.Background()
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://thanos:Ep0KmT1cvYU4xAAb@34.134.195.7:37017/"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Db Connected")
	return client.Database("ezyevents")
}
