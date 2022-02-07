package openweathermap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
)

var appid string

func Init(a string) {
	appid = a
}

type MainWeather struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int       `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int       `json:"timezone"`
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Type    int    `json:"type"`
	Id      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

func (w MainWeather) GetCelsius() float64 {
	return math.Round(w.Main.Temp - 273.15)
}

func (w MainWeather) GetCelsiusMin() float64 {
	return math.Round(w.Main.TempMin - 273.15)
}

func (w MainWeather) GetCelsiusMax() float64 {
	return math.Round(w.Main.TempMax - 273.15)
}

func (w MainWeather) GetClouds() string {
	if len(w.Weather) > 0 {
		return w.Weather[0].Description
	}
	return ""
}

func GetWeather(cityUser float64) (MainWeather, error) {
	var result MainWeather

	var cityNum float64 = cityUser
	var link = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?id=%.0f&appid=%s", cityNum, appid)
	fmt.Println("http.Get url: ", link)
	resp, err := http.Get(link)
	if err != nil {
		return result, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal([]byte(b), &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
