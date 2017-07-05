package main

import (
	"net/http"
	"fmt"
	"log"
	"os"
)

var (
	slackWebhook string
	weatherAPIKey string
)

func main() {
	port := os.Getenv("PORT")
	weatherAPIKey = os.Getenv("WEATHERAPIKEY")

	if port == "" {
		log.Fatal("$PORT must be set")
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