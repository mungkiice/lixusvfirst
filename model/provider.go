package model

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Provider struct {
	ID    primitive.ObjectID `bson:"_id, omitempty" json:"-"`
	Name  string             `bson:"name"`
	Code  []string           `bson:"codes"`
	Price float64            `bson:"price"`
}

func FindAllProviders(c *mongo.Client, filter bson.M) (result []*Provider) {
	col := c.Database("vfirst").Collection("provider")
	cursor, err := col.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on finding provider documents:", err)
	}
	for cursor.Next(context.TODO()) {
		var provider Provider
		err := cursor.Decode(&provider)
		if err != nil {
			log.Fatal("Error on decoding provider document:", err)
		}
		result = append(result, &provider)
	}
	return
}

func FindMatchProvider(c *mongo.Client, phone string) (result *Provider, exists bool) {
	regex, err := regexp.Compile(`8\d{2}`)
	if err != nil {
		log.Fatal("Error on compiling regex")
	}
	pattern := regex.FindString(phone)
	providers := FindAllProviders(c, bson.M{})
	for _, p := range providers {
		fmt.Println("Providers", p.Name)
		for _, code := range p.Code {
			fmt.Println("Provider Code:", code)
			if code == pattern {
				fmt.Println("Provider matched!")
				return p, true
			}
		}
	}
	fmt.Println("providers empty")
	return
}
