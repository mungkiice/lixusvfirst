package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

type SMSRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	To       string `json:"to"`
	From     string `json:"from"`
	Text     string `json:"text"`
}

func pushSMS(c *gin.Context) {
	var req SMSRequest
	var response string
	var err error
	var udh = ""
	if err = c.ShouldBind(&req); err != nil {
		log.Println("Error on binding user request:", err)
	}
	var dlrURL = "http://103.129.223.17:8080/app/status?" +
		"unique_id=%7&reason=%2&to=%p&from=%P&time=%t&status=%d" +
		"&delivered=%3&status_err=%4&client_guid=%5&client_sn=%6&" +
		"circle=%8&operator=%9&txt_status=%13&submit_date=%14&msg_status=%16"
	if len(req.Text) > 800 {
		var pivot, i = 0, 1
		rand.Seed(time.Now().UnixNano())
		refCode := fmt.Sprintf("%X", rand.Intn(255))
		smsCount := int(math.Ceil(float64(len(req.Text)) / 160))
		for len(req.Text[pivot:]) > 160 {
			udh = fmt.Sprintf("050003%v%02d%02d", refCode, smsCount, i)
			log.Println("Multiple SMS UDH: ", udh)
			response, err = sendReq(req.Username, req.Password, req.To, udh, req.From, req.Text[pivot:pivot+160], dlrURL)
			if err != nil {
				log.Printf("Error on sending SMS %d: %v\n", i, err)
			}
			pivot += 160
			i++
		}
		udh = fmt.Sprintf("050003%v%02d%02d", refCode, smsCount, i)
		log.Println("Multiple SMS UDH: ", udh)
		response, err = sendReq(req.Username, req.Password, req.To, udh, req.From, req.Text[pivot:len(req.Text)], dlrURL)
		if err != nil {
			log.Printf("Error on sending last SMS: %v\n", err)
		}
	} else {
		log.Println("Single SMS UDH: ", udh)
		response, err = sendReq(req.Username, req.Password, req.To, udh, req.From, req.Text, dlrURL)
		if err != nil {
			log.Printf("Error on sending request to VFirst: %v\n", err)
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}

func sendReq(uname, pass, to, udh, from, text, dlrURL string) (string, error) {
	url := fmt.Sprintf("http://www.myvaluefirst.com/smpp/sendsms?username=%s&password=%s&to=%s&udh=%s&from=%s&text=%s&dlr-url=%s",
		uname, pass, to, url.PathEscape(udh), from, url.PathEscape(text), url.QueryEscape(dlrURL))
	log.Println("Sending req with url:", url)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
