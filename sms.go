package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SMS struct {
	To              string    `bson:"to"`
	From            string    `bson:"from"`
	Message         string    `bson:"message"`
	Time            time.Time `bson:"time"`
	DeliveredDate   time.Time `bson:"delivered_date`
	ClientGUID      string    `bson:"client_guid"`
	ClientSeqNumber string    `bson:"client_seq_number"`
	MessageID       string    `bson:"message_id"`
	Circle          string    `bson:"circle"`
	Operator        string    `bson:"operator"`
	MSGStatus       string    `bson:"msg_status"`
	VendorStatus    string    `bson:"vendor_status"`
	Client          string    `bson:"client"`
	// TextStatus      string    `bson:"text_status"`
	// SubmitDate      time.Time `bson:"submit_date"`
	// MessageStatus   int       `bson:"message_status"`
	// ReasonCode      string    `bson:"reason_code"`
	// StatusError     string    `bson:"status_error"`
}

func findAllStatus(c *mongo.Client, filter bson.M) (result []*SMS) {
	collection := c.Database("vfirst").Collection("SMS")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Println("Error on finding the documents:", err)
	}
	for cur.Next(context.TODO()) {
		var status SMS
		err := cur.Decode(&status)
		if err != nil {
			log.Println("Error on decoding the document:", err)
		}
		result = append(result, &status)
	}
	return
}

// func findOneStatus(c *mongo.Client, filter bson.M) (result SMSStatus) {
// 	c.Database("vfirst").Collection("SMS").
// 		FindOne(context.TODO(), filter).Decode(&result)
// 	return
// }

func addSMS(c *mongo.Client, status SMS) interface{} {
	result, err := c.Database("vfirst").Collection("SMS").InsertOne(context.TODO(), status)
	if err != nil {
		log.Println("Error on inserting new document:", err)
	}
	return result.InsertedID
}

// func removeStatus(c *mongo.Client, filter bson.M) int64 {
// 	result, err := c.Database("vfirst").Collection("SMS").DeleteOne(context.TODO(), filter)
// 	if err != nil {
// 		log.Println("Error on deleting document:", err)
// 	}
// 	return result.DeletedCount
// }

func updateSMS(c *mongo.Client, filter bson.M, newData bson.M) int64 {
	result, err := c.Database("vfirst").Collection("SMS").UpdateOne(
		context.TODO(), filter, bson.D{{Key: "$set", Value: newData}})
	if err != nil {
		log.Println("Error on updating document:", err)
	}
	log.Println("updateSMS result", result.ModifiedCount)
	return result.ModifiedCount
}
