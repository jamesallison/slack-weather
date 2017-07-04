package main

import (
	"net/http"
	"fmt"
	"log"
	"os"

	"github.com/jamesallison/slack-weather/weather"
)

var (
	slackWebhook string
	weatherAPIKey string
)

func main() {
	port := os.Getenv("PORT")
	slackWebhook = os.Getenv("SLACKWEBHOOK")
	weatherAPIKey = os.Getenv("WEATHERAPIKEY")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	if slackWebhook == "" {
		// set to default
		slackWebhook = "https://hooks.slack.com/services/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"
	}
	if weatherAPIKey == "" {
		// set to default
		weatherAPIKey = "4a960eda5ddf27ae6b3c6a899d9c5e39"
	}

	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":" + port, nil)
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	log.Print(slackWebhook)
	fmt.Fprintf(w, "Hey 😄")
}