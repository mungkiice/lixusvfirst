package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SMS struct {
	ID              primitive.ObjectID `bson:"_id, omitempty"`
	To              string             `bson:"to"`
	From            string             `bson:"from"`
	Message         string             `bson:"message"`
	Time            *time.Time         `bson:"time"`
	DeliveredDate   *time.Time         `bson:"delivered_date`
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

func FindLatestMatchSMS(c *mongo.Client, filter bson.D) SMS {
	opt := options.Find()
	opt.SetSort(bson.D{{"_id", -1}})
	opt.SetLimit(1)
	cursor, err := c.Database("vfirst").Collection("SMS").
		Find(context.TODO(), filter, opt)
	if err != nil {
		log.Fatal("Error while finding latest sms:", err)
	}
	for cursor.Next(context.TODO()) {
		var sms SMS
		err := cursor.Decode(&sms)
		if err != nil {
			log.Fatal("Error while decoding latest sms:", err)
		}
		return sms
	}
}

func AddSMS(c *mongo.Client, status SMS) interface{} {
	result, err := c.Database("vfirst").Collection("sms").InsertOne(context.TODO(), status)
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

func (sms *SMS) UpdateSMS(c *mongo.Client, newData bson.M) int64 {
	result, err := c.Database("vfirst").Collection("sms").UpdateOne(
		context.TODO(), bson.M{"_id": sms.ID}, bson.D{{Key: "$set", Value: newData}})
	if err != nil {
		log.Println("Error on updating document:", err)
	}
	log.Println("updateSMS result", result.ModifiedCount)
	return result.ModifiedCount
}
