package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type SMSRequest struct {
	To   string `json:"to"		form:"to"`
	From string `json:"from"	form:"from"`
	Text string `json:"text"	form:"text"`
}

var mc = getClient()

func main() {
	err := mc.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	// f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }
	// defer f.Close()

	// log.SetOutput(f)
	// log.Println("This is a test log entry")

	router := gin.Default()
	router.GET("/app/status", saveStatus)
	router.GET("/app/list", listStatus)
	router.POST("/app/push", pushSMS)
	router.Run(":8080")
	// var dlrURL = "http://103.129.223.17:8080/app/status?" +
	// 	"unique_id=%7&reason=%2&to=%p&from=%P&time=%t&status=%d" +
	// 	"&delivered=%3&status_err=%4&client_guid=%5&client_sn=%6&" +
	// 	"circle=%8&operator=%9&txt_status=%13&submit_date=%14&msg_status=%16"
	// log.Println("DLR URL :", dlrURL)
	// log.Println("PathEscape DLR URL", url.PathEscape(dlrURL))
	// log.Println("QueryEscape DLR URL", url.QueryEscape(dlrURL))
}
