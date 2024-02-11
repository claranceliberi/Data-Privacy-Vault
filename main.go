package main

import (
	"log"
	"net/http"

	"github.com/claranceliberi/data-privacy-vault/db"
	"github.com/claranceliberi/data-privacy-vault/utils"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	type Data map[string]string

	type TokenizeRequest struct {
		Id   string `json:"id"`
		Data Data   `json:"data"`
	}

	r.POST("/tokenize", func(c *gin.Context) {

		var tokenizeRequest TokenizeRequest

		if err := c.ShouldBindJSON(&tokenizeRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		encrypted := make(map[string]string)

		for key, value := range tokenizeRequest.Data {
			encryptedValue := utils.Encrypt(value, utils.MySecret)
			// generate unique token for data
			token := utils.Tokenize(value)

			// store token in redis with data encrypted
			err := db.Client.Set(token, encryptedValue, 0).Err()

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}

			encrypted[key] = token
		}

		c.JSON(http.StatusOK, gin.H{"id": tokenizeRequest.Id, "data": encrypted})

	})

	r.POST("/detokenize", func(c *gin.Context) {

		var detokenizeRequest TokenizeRequest

		if err := c.ShouldBindJSON(&detokenizeRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		decrypted := make(map[string]map[string]interface{})

		for key, value := range detokenizeRequest.Data {
			solution := make(map[string]interface{})

			if val := db.Client.Get(value); val.Err() == nil {
				solution["found"] = true
				log.Println(val.Val())
				solution["value"] = utils.Decrypt(val.Val(), utils.MySecret)
			} else {
				solution["found"] = false
				solution["value"] = ""
			}

			decrypted[key] = solution
		}
		c.JSON(http.StatusOK, gin.H{"id": detokenizeRequest.Id, "data": decrypted})

	})

	return r

}

func main() {
	// initialize db
	db.Init()

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
