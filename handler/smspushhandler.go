package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"github.com/mungkiice/vfirst/config"
	"github.com/mungkiice/vfirst/database"
	"github.com/mungkiice/vfirst/model"
	"go.mongodb.org/mongo-driver/bson"
)

type smsRequest struct {
	To   string `json:"to"`
	From string `json:"from"`
	Text string `json:"text"`
}

func PushSMS(c *gin.Context) {
	var req smsRequest
	var response string
	var err error
	var client model.Client
	var udh = ""

	uname, ok := c.Get("uname")
	if !ok {
		fmt.Println("Error uname doesnt exists in context")
	}

	if err := model.FindOneClient(database.Conn, bson.M{"username": uname.(string)}, &client); err != nil {
		fmt.Println("Error on finding match client by username:", err)
	}

	if err = c.ShouldBind(&req); err != nil {
		fmt.Println("Error on binding user request:", err)
	}

	provider, exists := model.FindMatchProvider(database.Conn, req.To)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "match provider not found",
		})
		return
	}

	quota := reflect.Indirect(reflect.ValueOf(&client)).FieldByName(strings.Title(provider.Name))
	if quota.Int() < 1 {
		c.JSON(http.StatusOK, gin.H{
			"error": "your quota is used up " + provider.Name,
		})
		return
	}

	var dlrURL = "http://" + config.GetObject().Server.Host + ":" +
		config.GetObject().Server.Port + "/status?unique_id=%7&" +
		"reason=%2&to=%p&from=%P&time=%t&status=%d&delivered=" +
		"%3&status_err=%4&client_guid=%5&client_sn=%6&circle=" +
		"%8&operator=%9&txt_status=%13&submit_date=%14&msg_status=%16"

	if len(req.Text) > 800 {
		var pivot, i = 0, 1
		rand.Seed(time.Now().UnixNano())
		refCode := fmt.Sprintf("%X", rand.Intn(255))
		smsCount := int(math.Ceil(float64(len(req.Text)) / 160))
		for len(req.Text[pivot:]) > 160 {
			udh = fmt.Sprintf("050003%v%02d%02d", refCode, smsCount, i)
			fmt.Println("Multiple SMS UDH: ", udh)
			response, err = sendReq(client.Username, client.Pass, req.To, udh, req.From, req.Text[pivot:pivot+160], dlrURL)
			if err != nil {
				fmt.Printf("Error on sending SMS %d: %v\n", i, err)
			}
			pivot += 160
			i++
		}
		udh = fmt.Sprintf("050003%v%02d%02d", refCode, smsCount, i)
		fmt.Println("Multiple SMS UDH: ", udh)
		response, err = sendReq(client.Username, client.Pass, req.To, udh, req.From, req.Text[pivot:len(req.Text)], dlrURL)
		if err != nil {
			fmt.Printf("Error on sending last SMS: %v\n", err)
		}
	} else {
		fmt.Println("Single SMS UDH: ", udh)
		response, err = sendReq(client.Username, client.Pass, req.To, udh, req.From, req.Text, dlrURL)
		if err != nil {
			fmt.Printf("Error on sending request to VFirst: %v\n", err)
		}
	}

	model.AddSMS(database.Conn, bson.M{
		"_id":           primitive.NewObjectID(),
		"to":            req.To,
		"from":          req.From,
		"message":       req.Text,
		"vendor_status": response,
		"client":        client.Username,
	})

	c.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}

func sendReq(uname, pass, to, udh, from, text, dlrURL string) (string, error) {
	url := fmt.Sprintf("http://www.myvaluefirst.com/smpp/sendsms?username=%s&password=%s&to=%s&udh=%s&from=%s&text=%s&dlr-url=%s",
		uname, pass, to, url.PathEscape(udh), from, url.PathEscape(text), url.QueryEscape(dlrURL))
	fmt.Println("Sending req with url:", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error on calling myvalue endpoint")
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error on reading myvalue response")
		return "", err
	}
	return string(body), nil
}
