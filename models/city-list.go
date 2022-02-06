package models

import (
	"encoding/json"

	"fmt"
	"io/ioutil"
	"os"
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

	// Open our jsonFile
	jsonFile, err := os.Open("citylist.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		return result, err
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return result, err
	}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &result)
	return result, err

	// 1. новая функция должна быть за пределами другой функции
	// 2. когда функция принимает переменную, у нее должнен быть указан тип
	// 3. когда функция ожидает возвращения какой то переменной, ее нужно вернуть return-ом
	// 4. если ты называешь функцию getCity то логично чтобы она возвращала City

}

func (c CityList) GetCityId(cityName string) (float64, error) {
	for _, city := range c {
		if cityName == city.Name {
			return city.Id, nil
		}
	}
	return -1, fmt.Errorf("unknown city")
}
