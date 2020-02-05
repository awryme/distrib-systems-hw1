package main

import (
	"context"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		panic(err)
	}
	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	const dbName = "incrementer"
	collection := mongoClient.Database(dbName).Collection(dbName)

	handlers := &Handlers{mongoCollection:collection}
	// Routes
	api := e.Group("/api")
	v1 := api.Group("/v1")
	v1.POST("/inc", handlers.incrementerHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}