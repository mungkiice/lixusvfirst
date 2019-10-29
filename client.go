package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	Username string `bson:"username" json="username"`
	Pass     string `bson:"password" json="-"`
	Token    string `bson:"token" json="token"`
}

func findOneClient(c *mongo.Client, filter bson.M, client *Client) error {
	return c.Database("vfirst").Collection("client").
		FindOne(context.TODO(), filter).Decode(client)

}
