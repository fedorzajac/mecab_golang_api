package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"

	"main/categorize"
)

var t *tokenizer.Tokenizer
var apiKey = "fedorfedor"

func init() {
	// initialize tokenizer
	var err error
	t, err = tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		panic(err)
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.Header.Get("X-API-Key")
		if key != apiKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			c.Abort()
			return
		}
		c.Next()
	}
}



type message struct {
	Surface string
	Features string
	Category categorize.Cat
}

func MecabHandler(c *gin.Context) {
	// parse json post parameters
	var json struct {
		Text string `json:"text"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// tokenize
	tokens := t.Tokenize(json.Text)

	// format result as json
	var result []message
	for _, token := range tokens {
		features := strings.Join(token.Features(), ",")
		result = append(result, message{
			Surface:  token.Surface,
			Features: features,
			Category: categorize.Categorize(token.Surface),
		})
	}
	c.JSON(http.StatusOK, result)
}

func main() {
	router := gin.Default()

	router.Use(authMiddleware())

	router.POST("/", MecabHandler)

	if err := router.Run("0.0.0.0:8080"); err != nil {
		fmt.Println(err)
	}
}
