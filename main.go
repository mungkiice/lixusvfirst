package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	// f, err := os.OpenFile("/var/www/vfirst/vfirst.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
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
	f, _ := os.Create("/usr/share/nginx/html/gin.log")
	gin.DefaultWriter = f

	router := gin.Default()
	r := multitemplate.NewRenderer()
	r.AddFromFiles("dashboard_page", "/usr/share/nginx/html/index.html")
	router.HTMLRender = r
	router.Static("/public", "/usr/share/nginx/html/public")
	router.GET("/", dashboard)
	router.GET("/status", updateStatus)
	router.GET("/list", VerifyToken(), listSMS)
	router.POST("/push", VerifyToken(), pushSMS)
	router.POST("/login", doLogin)
	router.Run(":8080")
}
