package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func dashboard(c *gin.Context) {
	sms := findAllSMS(mc, bson.M{})
	fmt.Println(sms)
	c.HTML(http.StatusOK, "dashboard_page", gin.H{
		"sms": sms,
	})
}
