package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mungkiice/vfirst/database"
	"github.com/mungkiice/vfirst/model"
	"go.mongodb.org/mongo-driver/bson"
)

func ShowDashboard(c *gin.Context) {
	sms := model.FindAllSMS(database.Conn, bson.M{})
	c.HTML(http.StatusOK, "dashboard_page", gin.H{
		"sms": sms,
	})
}
