package main

import (
	"encoding/json"
	"fmt"
	"fpgo/reduce/utils"
	"os"
)

type Entry struct {
	Airport struct {
		Code string `json:"Code"`
		Name string `json:"Name"`
	} `json:"Airport"`
	Statistics struct {
		Flights struct {
			Cancelled int `json:"Cancelled"`
			Delayed   int `json:"Delayed"`
			OnTime    int `json:"On Time"`
			Total     int `json:"Total"`
		} `json:"Flight"`
		MinutesDelayed struct {
			Carrier      int `json:"Carrier"`
			LateAircraft int `json:"Late Aircraft"`
			Security     int `json:"Security"`
			Weather      int `json:"Weather"`
		} `json:"Minutes Delayed"`
	} `json:"Statistics"`
}

func getEnteries() []Entry {
	bytes, err := os.ReadFile("./data/airlines.json")
	if err != nil {
		panic(err)
	}
	var entries []Entry
	err = json.Unmarshal(bytes, &entries)
	if err != nil {
		panic(err)
	}
	return entries
}

func main() {
	// Sum
	ints := []int{1, 2, 3, 4}
	sum := utils.Sum(ints)
	fmt.Printf("%v\n", sum)

	// Product
	product := utils.Product(ints)
	fmt.Printf("%v\n", product)

	//
	words := []string{"hello", "world", "universe"}
	result := utils.ReduceWithStart(words, "first", func(s1, s2 string) string {
		return s1 + ", " + s2
	})
	fmt.Printf("%v\n", result)

	//
	entries := getEnteries()
	SEA := utils.Filter(entries, func(e Entry) bool {
		return e.Airport.Code == "SEA"
	})
	WeatherDelayHours := utils.FMap(SEA, func(e Entry) int {
		return e.Statistics.MinutesDelayed.Weather / 60
	})
	totalWeatherDelay := utils.Sum(WeatherDelayHours)
	fmt.Printf("%v\n", totalWeatherDelay)

}
