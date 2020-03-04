package mongostorage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"incrementer/storage"
)

const queryDbKey = "query"

type MongoStorage struct {
	Collection *mongo.Collection
}

var _ storage.Storage = (*MongoStorage)(nil)

func (storage *MongoStorage) HasEither(ctx context.Context, first, second int) (bool, bool, error) {
	cursor, err := storage.Collection.Find(ctx, bson.M{
		"$or": []interface{}{
			bson.M{queryDbKey: first},
			bson.M{queryDbKey: second},
		},
	})
	if err == mongo.ErrNoDocuments {
		return false, false, nil
	}
	if err != nil {
		return false, false, fmt.Errorf("failed to find in mongo: %w", err)
	}
	defer cursor.Close(ctx)
	allRes := make([]bson.M, 2)
	err = cursor.All(ctx, &allRes)
	if err != nil {
		return false, false, fmt.Errorf("failed to iterate cursor: %w", err)
	}
	for _, dbRes := range allRes {
		q := int(dbRes[queryDbKey].(int32))
		if q == first {
			return true, false, nil
		}
		if q == second {
			return false, true, nil
		}
	}
	return false, false, nil
}

func (storage *MongoStorage) Insert(ctx context.Context, number int) error {
	_, err := storage.Collection.InsertOne(ctx, bson.M{
		queryDbKey: number,
	})
	if err != nil {
		return fmt.Errorf("failed to insert '%d': %w", number, err)
	}
	return nil
}
