package model

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	ID        primitive.ObjectID `bson:"_id, omitempty" json:"-"`
	Username  string             `bson:"username" json:"username"`
	Pass      string             `bson:"password" json:"-"`
	Token     string             `bson:"token" json:"token"`
	Telkomsel int64              `bson:"telkomsel"`
	Xl        int64              `bson:"xl"`
	Tri       int64              `bson:"tri"`
	Indosat   int64              `bson:"indosat"`
}

func FindOneClient(c *mongo.Client, filter bson.M, client *Client) error {
	return c.Database("vfirst").Collection("client").
		FindOne(context.TODO(), filter).Decode(client)
}

func (cl *Client) DeductQuota(c *mongo.Client, providerName string) {
	_, err := c.Database("vfirst").Collection("client").
		UpdateOne(context.TODO(), bson.M{"_id": cl.ID},
			bson.D{{Key: "$inc", Value: bson.M{providerName: -1}}})
	if err != nil {
		log.Fatal("Error on deducting quota ", providerName, err)
	}
}
