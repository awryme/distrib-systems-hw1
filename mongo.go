package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func getMongoCollection(addr, dbName, collectionName string) (*mongo.Collection, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	mongoClient, err := mongo.Connect(ctx,
		options.Client().ApplyURI(addr),
		options.Client().SetMaxPoolSize(10),
	)
	if err != nil {
		return nil, fmt.Errorf("failed get mongo connection: %w", err)
	}

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping mongo: %w", err)
	}
	collection := mongoClient.Database(dbName).Collection(collectionName)
	return collection, nil
}
