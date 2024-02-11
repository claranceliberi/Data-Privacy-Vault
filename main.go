package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"
var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}


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

		requestId := tokenizeRequest.Id
		// requestData := tokenizeRequest.Data



		c.JSON(http.StatusOK, gin.H{
			"id": encrypt(requestId, MySecret),
		})

	})

	r.POST("/detokenize", func(c *gin.Context) {

	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
   }
   
   

func encrypt(data string, key string) string {
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		panic(err)
	}

	plaintext := []byte(data)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	cipherText := make([]byte, len(plaintext))
	cfb.XORKeyStream(cipherText, plaintext)

	return Encode(cipherText)
}

func decrypt(data string, key string) string {
	return "string"
}