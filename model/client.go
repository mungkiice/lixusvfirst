package model

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	ID       primitive.ObjectID `bson:"_id, omitempty" json:"-"`
	Username string             `bson:"username" json:"username"`
	Pass     string             `bson:"password" json:"-"`
	Balance  float64            `bson:"balance" json:"balance"`
	Token    string             `bson:"token" json:"token"`
}

func FindOneClient(c *mongo.Client, filter bson.M, client *Client) error {
	return c.Database("vfirst").Collection("client").
		FindOne(context.TODO(), filter).Decode(client)
}

func (cl *Client) PayBill(c *mongo.Client, amount float64) {
	newBalance := cl.Balance - amount
	_, err := c.Database("vfirst").Collection("client").
		UpdateOne(context.TODO(), bson.M{"_id": cl.ID},
			bson.D{{Key: "$set", Value: bson.M{"balance": newBalance}}})
	if err != nil {
		log.Fatal("Error on paying bill:", err)
	}
}
