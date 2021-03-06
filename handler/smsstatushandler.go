package handler

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mungkiice/vfirst/database"
	"github.com/mungkiice/vfirst/model"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	timeLayout = "2006-01-02 15:04:05"
)

func ListSMS(c *gin.Context) {
	uname, _ := c.Get("uname")
	c.JSON(http.StatusOK, gin.H{
		"SMS Status List": model.FindAllSMS(database.Conn, bson.M{"client": uname.(string)}),
	})
}

func UpdateStatus(c *gin.Context) {
	fmt.Println(c.Request.URL.Query())
	uniqueID := c.Query("unique_id")
	receiver := c.Query("to")
	sender := c.Query("from")
	responseTime := c.Query("time")
	delivered := c.Query("delivered")
	clientGUID := c.Query("client_guid")
	clientSN := c.Query("client_sn")
	circle := c.Query("circle")
	operator := c.Query("operator")
	msgStatus := c.Query("msg_status")
	// textStatus := c.Query("txt_status")
	// submit := c.Query("submit_date")
	// statusErr := c.Query("status_err")
	// status := c.Query("status")
	// reasonCode := c.Query("reason")
	messageTime, err := time.Parse(timeLayout, responseTime)
	if err != nil {
		log.Fatal("Error on converting message time:", err)
	}
	// messageStatus, err := strconv.Atoi(status)
	// if err != nil {
	// 	log.Println("Error on converting message status:", err)
	// }
	deliveredDate, err := time.Parse(timeLayout, delivered)
	if err != nil {
		log.Fatal("Error on converting delivered date:", err)
	}
	// submitDate, err := time.Parse(timeLayout, submit)
	// if err != nil {
	// 	log.Println("Error on converting submit date:", err)
	// }
	regex, _ := regexp.Compile(`[A-Z]+\d+`)
	// sort, err := options.Sort(bson.NewDocument(bson.EC.Int32("word", 1)))
	// if err != nil {
	// 	log.Println("Error on sort")
	// }
	clientUname := strings.ToLower(regex.FindString(clientGUID))
	sms, exists := model.FindLatestMatchSMS(database.Conn, bson.D{
		{"to", receiver},
		{"from", sender},
		{"client", clientUname},
	})
	if exists {
		provider, exists := model.FindMatchProvider(database.Conn, receiver)
		if exists {
			var client model.Client
			if err := model.FindOneClient(database.Conn, bson.M{"username": clientUname}, &client); err != nil {
				fmt.Println("Error on finding match client by username:", err)
			}
			client.DeductQuota(database.Conn, provider.Name)
			sms.UpdateSMS(database.Conn, bson.M{
				"delivered_date":    deliveredDate,
				"client_guid":       clientGUID,
				"message_id":        uniqueID,
				"client_seq_number": clientSN,
				"circle":            circle,
				"msg_status":        msgStatus,
				"operator":          operator,
				"time":              messageTime,
			})
		}
	}
}
