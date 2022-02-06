package models

import (
	"math"
)

type MainWeather struct {
	Coord      Coord
	Weather    []Weather
	Base       string `json:"base"`
	Main       Main
	Visibility int `json:"visibility"`
	Wind       Wind
	Clouds     Clouds
	Dt         int `json:"dt"`
	Sys        Sys
	Timezone   int    `json:"timezone"`
	Id         int    `json:"703448"`
	Name       string `json:"name"`
	Cod        int    `json:"cod"`
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
	Temp       float64 `json:"temp"`
	Feels_like float64 `json:"feels_like"`
	Temp_min   float64 `json:"temp_min"`
	Temp_max   float64 `json:"temp_max"`
	Pressure   int     `json:"pressure"`
	Humidity   int     `json:"humidity"`
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
