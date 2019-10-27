package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	// f, err := os.OpenFile("/var/www/vfirst.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Println("error opening file:", err)
	// }
	// defer f.Close()

	// log.SetOutput(f)
	// log.Println("This is a test log entry")

	if err := mc.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Println("Couldn't connect to the database", err)
	} else {
		log.Println("Database connected!")
	}
}

func main() {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("/var/www/gin.log")
	gin.DefaultWriter = f

	router := gin.Default()
	router.GET("/app/status", updateStatus)
	router.GET("/app/list", VerifyToken(), listStatus)
	router.POST("/app/push", VerifyToken(), pushSMS)
	router.POST("/app/login", doLogin)
	router.GET("/app/last", func(c *gin.Context) {
		sms := findOneSMS(mc, bson.M{})
		c.JSON(http.StatusOK, gin.H{
			"sms": sms,
		})

	})
	router.Run(":8080")
}
