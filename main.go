package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/gin-gonic/gin"
)

func main() {
	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", "", "Twitter ConsumerSecret")
	accessToken := flags.String("access-token", "", "Twitter AccessToken")
	accessSecret := flags.String("access-secret", "", "Twitter AccessSecret")
	flags.Parse(os.Args[1:])

	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	r := gin.Default()

	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)
	// Twitter Client
	client := twitter.NewClient(httpClient)

	r.GET("/search", func(c *gin.Context) {
		search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
			Query: c.Query("q"),
		})
		if err != nil {
			fmt.Printf("error search %v\n", err)
		}
		c.JSON(http.StatusOK, search)
	})

	r.Static("/page", "./page")

	r.Run(":8088")
}
