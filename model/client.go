package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	ID       primitive.ObjectID `bson:"_id, omitempty" json:"-"`
	Username string             `bson:"username" json:"username"`
	Pass     string             `bson:"password" json:"-"`
	Bill     int64              `bson:"bill" json:"bill"`
	Token    string             `bson:"token" json:"token"`
}

func FindOneClient(c *mongo.Client, filter bson.M, client *Client) error {
	return c.Database("vfirst").Collection("client").
		FindOne(context.TODO(), filter).Decode(client)

}
