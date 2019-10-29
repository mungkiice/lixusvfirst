package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Provider struct {
	ID    primitive.ObjectID `bson:"_id, omitempty" json:"-"`
	Name  string             `bson:"name"`
	Price int64              `bson:"price"`
}
