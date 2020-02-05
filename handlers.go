package main

import (
	"fmt"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type Handlers struct {
	mongoCollection *mongo.Collection
}

func (h *Handlers) incrementerHandler(c echo.Context) error {
	ctx := c.Request().Context()
	req := &Request{}
	res := &Response{}

	err := c.Bind(req)
	if err != nil {
		res.Err = "failed to parse request: " + err.Error()
		c.Logger().Error(res.Err)
		return c.JSON(http.StatusBadRequest, res)
	}
	res.Query = req.Query
	res.Resp = req.Query + 1
	cursor, err := h.mongoCollection.Find(ctx, bson.M{
		"$or": []interface {}{
			bson.M{"query": req.Query+1},
			bson.M{"query": req.Query},
		},
	})
	if err == mongo.ErrNoDocuments {
		res.Err = ""
		_, err := h.mongoCollection.InsertOne(ctx, bson.M{
			"query": req.Query,
		})
		fmt.Println("inserted [no docs]", req.Query, err)
		if err != nil {
			res.Err = "internal error: " + err.Error()
			c.Logger().Error(res.Err)
			return c.JSON(http.StatusInternalServerError, res)
		}
		return c.JSON(http.StatusOK, res)
	}
	if err != nil {
		res.Err = "internal error: " + err.Error()
		c.Logger().Error(res.Err)
		return c.JSON(http.StatusInternalServerError, res)
	}
	defer cursor.Close(ctx)
	allRes := make([]bson.M, 0, 10)
	err = cursor.All(ctx, &allRes)
	if err != nil {
		res.Err = "internal error: " + err.Error()
		c.Logger().Error(res.Err)
		return c.JSON(http.StatusInternalServerError, res)
	}
	for _, dbRes := range allRes {
		fmt.Println("dbres", dbRes)
		q := int(dbRes["query"].(int32))
		if q == req.Query {
			res.Err = "already done this number"
			c.Logger().Error(res.Err)
			return c.JSON(http.StatusBadRequest, res)
		}
		if q == req.Query + 1 {
			res.Err = "already done this number+1"
			c.Logger().Error(res.Err)
			return c.JSON(http.StatusBadRequest, res)
		}
	}
	res2, err := h.mongoCollection.InsertOne(ctx, bson.M{
		"query": req.Query,
	})
	fmt.Println("inserted", req.Query, res2.InsertedID,  err)
	if err != nil {
		res.Err = "internal error: " + err.Error()
		c.Logger().Error(res.Err)
		return c.JSON(http.StatusInternalServerError, res)
	}
	return c.JSON(http.StatusOK, res)
}
