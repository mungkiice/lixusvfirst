package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SMSStatus struct {
	To              string    `bson:"to"`
	From            string    `bson:"from"`
	Time            time.Time `bson:"time"`
	MessageStatus   int       `bson:"message_status"`
	ReasonCode      string    `bson:"reason_code"`
	DeliveredDate   time.Time `bson:"delivered_date`
	StatusError     string    `bson:"status_error"`
	ClientGUID      string    `bson:"client_guid"`
	ClientSeqNumber string    `bson:"client_seq_number"`
	MessageID       string    `bson:"message_id"`
	Circle          string    `bson:"circle"`
	Operator        string    `bson:"operator"`
	TextStatus      string    `bson:"text_status"`
	SubmitDate      time.Time `bson:"submit_date"`
	MSGStatus       string    `bson:"msg_status"`
}

func findAllStatus(c *mongo.Client, filter bson.M) (result []*SMSStatus) {
	collection := c.Database("vfirst").Collection("SMS")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on finding the documents", err)
	}
	for cur.Next(context.TODO()) {
		var status SMSStatus
		err := cur.Decode(&status)
		if err != nil {
			log.Fatal("Error on decoding the document", err)
		}
		result = append(result, &status)
	}
	return
}

func findOneStatus(c *mongo.Client, filter bson.M) (result SMSStatus) {
	c.Database("vfirst").Collection("SMS").
		FindOne(context.TODO(), filter).Decode(&result)
	return
}

func addStatus(c *mongo.Client, status SMSStatus) interface{} {
	result, err := c.Database("vfirst").Collection("SMS").InsertOne(context.TODO(), status)
	if err != nil {
		log.Fatal("Error on inserting new document", err)
	}
	return result.InsertedID
}

func removeStatus(c *mongo.Client, filter bson.M) int64 {
	result, err := c.Database("vfirst").Collection("SMS").DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on deleting document", err)
	}
	return result.DeletedCount
}

func updateStatus(c *mongo.Client, filter bson.M, newData bson.M) int64 {
	result, err := c.Database("vfirst").Collection("SMS").UpdateOne(context.TODO(), filter, newData)
	if err != nil {
		log.Fatal("Error on updating document", err)
	}
	return result.ModifiedCount
}
