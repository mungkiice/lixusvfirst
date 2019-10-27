package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	f, err := os.OpenFile("vfirst.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println("error opening file:", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")

	if err := mc.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Println("Couldn't connect to the database", err)
	} else {
		log.Println("Database connected!")
	}
}

func main() {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = f

	router := gin.Default()
	router.GET("/app/status", saveStatus)
	router.GET("/app/list", listStatus)
	router.POST("/app/push", pushSMS)
	router.Run(":8080")
}
