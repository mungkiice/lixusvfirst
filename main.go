package main

import (
	"context"
	"log"
	"os"

	"github.com/mungkiice/vfirst/database"
	"github.com/mungkiice/vfirst/handler"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/mungkiice/vfirst/config"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	f, err := os.OpenFile("./vfirst.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println("error opening file:", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")

	if err := database.Conn.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Println("Couldn't connect to the database", err)
	} else {
		log.Println("Database connected!")
	}
}

func main() {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("./gin.log")
	gin.DefaultWriter = f

	router := gin.Default()
	r := multitemplate.NewRenderer()
	r.AddFromFiles("dashboard_page", "./web/index.html")
	router.HTMLRender = r
	router.Static("/public", "./public")
	router.GET("/", handler.ShowDashboard)
	router.GET("/status", handler.UpdateStatus)
	router.GET("/list", handler.VerifyToken(), handler.ListSMS)
	router.POST("/push", handler.VerifyToken(), handler.PushSMS)
	router.POST("/login", handler.DoLogin)
	router.Run(":" + config.GetObject().Server.Port)
}
