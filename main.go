package main

import (
	"net/http"
	"fmt"
	"log"
	"os"
)

var slackWebhook string

func main() {
	port := os.Getenv("PORT")
	slackWebhook = os.Getenv("SLACKWEBHOOK")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	if slackWebhook == "" {
		// set to default
		slackWebhook = "https://hooks.slack.com/services/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"
	}

	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":" + port, nil)
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	log.Print(slackWebhook)
	fmt.Fprintf(w, "Hey 😄")
}