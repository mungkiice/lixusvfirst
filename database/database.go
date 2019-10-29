package database

import (
	"context"
	"fmt"
	"log"

	"github.com/mungkiice/vfirst/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Conn *mongo.Client

func init() {
	Conn = GetClient()
}

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("%s://%s:%s",
		"mongodb", config.GetObject().Database.Host, config.GetObject().Database.Port))
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
