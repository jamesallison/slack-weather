package weather

import (
	"errors"
	"net/http"
	"encoding/json"
)

type Report struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID int `json:"id"`
		Main string `json:"main"`
		Description string `json:"description"`
		Icon string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp float64 `json:"temp"`
		Pressure int `json:"pressure"`
		Humidity int `json:"humidity"`
		TempMin float64 `json:"temp_min"`
		TempMax float64 `json:"temp_max"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg int `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt int `json:"dt"`
	Sys struct {
		Type int `json:"type"`
		ID int `json:"id"`
		Message float64 `json:"message"`
		Country string `json:"country"`
		Sunrise int `json:"sunrise"`
		Sunset int `json:"sunset"`
	} `json:"sys"`
	ID int `json:"id"`
	Name string `json:"name"`
	Cod int `json:"cod"`
}

func GetWeather(location, token string) (Report, error) {
	url := "http://api.openweathermap.org/data/2.5/weather?q=" + location + "&appid=" + token
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Report{}, err
	}

	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Report{}, err
	}
	if res.StatusCode != 200 {
		return Report{}, errors.New("Non-200 response code")
	}
	
	// parse json
	decoder := json.NewDecoder(res.Body)
	var weatherReport Report
	err = decoder.Decode(&weatherReport)
	if err != nil {
		return Report{}, err
	}

	return weatherReport, nil
}