package model

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SMS struct {
	ID              primitive.ObjectID `bson:"_id"`
	To              string             `bson:"to"`
	From            string             `bson:"from"`
	Message         string             `bson:"message"`
	Time            *time.Time         `bson:"time"`
	DeliveredDate   *time.Time         `bson:"delivered_date"`
	ClientGUID      string             `bson:"client_guid"`
	ClientSeqNumber string             `bson:"client_seq_number"`
	MessageID       string             `bson:"message_id"`
	Circle          string             `bson:"circle"`
	Operator        string             `bson:"operator"`
	MSGStatus       string             `bson:"msg_status"`
	VendorStatus    string             `bson:"vendor_status"`
	Client          string             `bson:"client"`
	// TextStatus      string    `bson:"text_status"`
	// SubmitDate      time.Time `bson:"submit_date"`
	// MessageStatus   int       `bson:"message_status"`
	// ReasonCode      string    `bson:"reason_code"`
	// StatusError     string    `bson:"status_error"`
}

func FindAllSMS(c *mongo.Client, filter bson.M) (result []*SMS) {
	collection := c.Database("vfirst").Collection("sms")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on finding sms documents:", err)
	}
	for cur.Next(context.TODO()) {
		var status SMS
		err := cur.Decode(&status)
		if err != nil {
			log.Fatal("Error on decoding sms document:", err)
		}
		result = append(result, &status)
	}
	return
}

func FindLatestMatchSMS(c *mongo.Client, filter bson.D) (sms SMS, exists bool) {
	fmt.Println("FindLatestMatchSMS")
	opt := options.Find()
	opt.SetSort(bson.D{{"_id", -1}})
	opt.SetLimit(1)
	cursor, err := c.Database("vfirst").Collection("sms").
		Find(context.TODO(), filter, opt)
	if err != nil {
		log.Fatal("Error while finding latest sms:", err)
	}
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&sms)
		if err != nil {
			log.Fatal("Error while decoding latest sms:", err)
		}
		fmt.Println("Found object with ID:", sms.ID)
		return sms, true
	}
	return SMS{}, false
}

func AddSMS(c *mongo.Client, sms bson.M) interface{} {
	fmt.Println("Adding sms to database")
	result, err := c.Database("vfirst").Collection("sms").InsertOne(context.TODO(), sms)
	if err != nil {
		fmt.Println("Error on inserting new document:", err)
	}
	fmt.Println("Result adding sms to database:", result.InsertedID)
	return result.InsertedID
}

// func removeStatus(c *mongo.Client, filter bson.M) int64 {
// 	result, err := c.Database("vfirst").Collection("SMS").DeleteOne(context.TODO(), filter)
// 	if err != nil {
// 		log.Println("Error on deleting document:", err)
// 	}
// 	return result.DeletedCount
// }

func (sms *SMS) UpdateSMS(c *mongo.Client, newData bson.M) int64 {
	fmt.Println("On update sms looking for ID:", sms.ID)
	result, err := c.Database("vfirst").Collection("sms").UpdateOne(
		context.TODO(), bson.M{"_id": sms.ID}, bson.D{{Key: "$set", Value: newData}})
	if err != nil {
		log.Fatal("Error on updating document:", err)
	}
	fmt.Println("updateSMS result", result.ModifiedCount)
	return result.ModifiedCount
}
