package main

import (
	"go.uber.org/zap"
	"incrementer/application"
	"incrementer/log"
	"incrementer/server"
	"incrementer/storage/mongostorage"
	"net/http"
)

const defaultMongoDb = "incrementer"
const defaultAddress = ":8000"

func main() {
	var mongoAddr = getEnvPanic("MONGO_ADDR")
	var mongoDb = getEnv("MONGO_DB", defaultMongoDb)
	var mongoCollection = getEnv("MONGO_COLLECTION", defaultMongoDb)
	var appAddress = getEnv("APP_ADDRESS", defaultAddress)

	collection, err := getMongoCollection(mongoAddr, mongoDb, mongoCollection)
	if err != nil {
		panic(err)
	}

	storage := &mongostorage.MongoStorage{
		Collection: collection,
	}

	app := &application.Application{
		Storage: storage,
	}

	var logger = log.New().With(zap.String("app", "incrementer"))

	handler := server.NewServer(logger, app)

	if err := http.ListenAndServe(appAddress, handler); err != nil {
		panic(err)
	}
}
