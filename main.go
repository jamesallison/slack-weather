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
	// see which location the user entered
	location := req.Form.Get("text")
	if location == "" {
		// default to London
		location = "London"
	}

	// get the weather for this location
	report, err := weather.GetWeather(location, weatherAPIKey)

	if err != nil {
		fmt.Fprintf(w, "Could not get the weather 😔")
		log.Print(err)
		return // the rest of the function will not run
	}

	fmt.Println(report.Name)

	// construct our lovely message
	message := "The weather in " + report.Name  + " is " + report.Weather[0].Description

	log.Print(slackWebhook)
	fmt.Fprintf(w, message)
}