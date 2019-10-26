package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	timeLayout = "2006-01-02 15:04:05"
)

func listStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"SMS Status List": findAllStatus(mc, bson.M{}),
	})
}

func saveStatus(c *gin.Context) {
	uniqueID := c.Query("unique_id")
	reasonCode := c.Query("reason")
	receiver := c.Query("to")
	sender := c.Query("from")
	responseTime := c.Query("time")
	status := c.Query("status")
	delivered := c.Query("delivered")
	statusErr := c.Query("status_err")
	clientGUID := c.Query("client_guid")
	clientSN := c.Query("client_sn")
	circle := c.Query("circle")
	operator := c.Query("operator")
	textStatus := c.Query("txt_status")
	submit := c.Query("submit_date")
	msgStatus := c.Query("msg_status")
	messageTime, err := time.Parse(timeLayout, responseTime)
	if err != nil {
		log.Fatal("Error on converting message time ", err)
	}
	messageStatus, err := strconv.Atoi(status)
	if err != nil {
		log.Fatal("Error on converting message status ", err)
	}
	deliveredDate, err := time.Parse(timeLayout, delivered)
	if err != nil {
		log.Fatal("Error on converting delivered date ", err)
	}
	submitDate, err := time.Parse(timeLayout, submit)
	if err != nil {
		log.Fatal("Error on converting submit date ", err)
	}
	var newStatus = SMSStatus{
		To:              receiver,
		From:            sender,
		Time:            messageTime,
		MessageStatus:   messageStatus,
		ReasonCode:      reasonCode,
		DeliveredDate:   deliveredDate,
		StatusError:     statusErr,
		ClientGUID:      clientGUID,
		ClientSeqNumber: clientSN,
		MessageID:       uniqueID,
		Circle:          circle,
		Operator:        operator,
		TextStatus:      textStatus,
		SubmitDate:      submitDate,
		MSGStatus:       msgStatus,
	}
	c.JSON(http.StatusOK, gin.H{
		"insertedID": addStatus(mc, newStatus),
	})
}
