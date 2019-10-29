package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mc *mongo.Client

func getClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("%s://%s:%s",
		"mongodb", getConfiguration().Database.Host, getConfiguration().Database.Port))
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func init() {
	mc = getClient()
}
