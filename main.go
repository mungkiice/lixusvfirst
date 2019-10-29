package main

import (
	"context"
	"log"
	"os"

	"github.com/mungkiice/vfirst/database"
	"github.com/mungkiice/vfirst/handler"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	f, err := os.OpenFile("/usr/share/nginx/html/vfirst.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("error opening file:", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")

	if err := database.Conn.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Database connected!")
	}
}

func main() {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("/usr/share/nginx/html/gin.log")
	gin.DefaultWriter = f

	router := gin.Default()
	r := multitemplate.NewRenderer()
	r.AddFromFiles("dashboard_page", "/usr/share/nginx/html/index.html")
	router.HTMLRender = r
	router.Static("/public", "/usr/share/nginx/html/public")
	router.GET("/", handler.ShowDashboard)
	router.GET("/status", handler.UpdateStatus)
	router.GET("/list", handler.VerifyToken(), handler.ListSMS)
	router.POST("/push", handler.VerifyToken(), handler.PushSMS)
	router.POST("/login", handler.DoLogin)
	router.Run(":8080")
}
