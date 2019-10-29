package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/mungkiice/vfirst/config"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mungkiice/vfirst/database"
	"github.com/mungkiice/vfirst/model"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Pass     string `json:"password" binding:"required"`
}

func DoLogin(c *gin.Context) {
	var client model.Client
	var req loginRequest

	if err := c.ShouldBind(&req); err != nil {
		log.Fatal("Error on binding user request:", err)
	}
	if err := model.FindOneClient(database.Conn, bson.M{"username": req.Username, "password": req.Pass}, &client); err != nil {
		log.Fatal("Error on finding match client:", err)
		c.JSON(http.StatusOK, gin.H{
			"message": "client credential invalid",
		})
		return
	}
	token, err := generateToken(&client)
	if err != nil {
		log.Fatal("Error on generating token:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   "login succeed",
		"token":    token,
		"username": client.Username,
	})
}

func generateToken(client *model.Client) (string, error) {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uname": client.Username,
		"exp":   time.Now().Add(time.Hour * 8760).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.GetObject().App.Key))

	return tokenString, err
}

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// sample token string taken from the New example
		tokenHeader := c.Request.Header.Get("Authorization")

		if !strings.Contains(tokenHeader, "Bearer") {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid token : method invalid",
			})
			c.Abort()
			return
		}

		tokenString := strings.Replace(tokenHeader, "Bearer ", "", -1)

		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(config.GetObject().App.Key), nil
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid token",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid token : claim invalid",
			})
			c.Abort()
			return
		}

		if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid token : token expired",
			})
			c.Abort()
			return
		}

		c.Set("uname", claims["uname"])
		c.Next()
	}
}
