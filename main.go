package main

import (
	"context"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoAddr = os.Getenv("MONGO_ADDR")

func main() {
	if mongoAddr == "" {
		panic("empty MONGO_ADDR")
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoAddr))
	if err != nil {
		panic(err)
	}
	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	const dbName = "incrementer"
	collection := mongoClient.Database(dbName).Collection(dbName)

	handlers := &Handlers{mongoCollection: collection}
	// Routes
	api := e.Group("/api")
	v1 := api.Group("/v1")
	v1.POST("/inc", handlers.incrementerHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
