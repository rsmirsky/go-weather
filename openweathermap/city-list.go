package openweathermap

import (
	"encoding/json"

	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type CityList []City

type City struct {
	Id      float64 `json:"id"`
	Name    string  `json:"name"`
	State   string  `json:"state"`
	Country string  `json:"country"`
	Coord   CoordCity
}

type CoordCity struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

func GetCityList() (CityList, error) {
	var result CityList

	jsonFile, err := os.Open("citylist.json")
	if err != nil {
		return result, err
	}

	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(byteValue, &result)
	return result, err
}

func (c CityList) GetCityId(cityName string) (float64, error) {
	userCityNameLower := strings.ToLower(cityName)
	for _, city := range c {
		if userCityNameLower == strings.ToLower(city.Name) {
			return city.Id, nil
		}
	}
	return -1, fmt.Errorf("unknown city")
}
