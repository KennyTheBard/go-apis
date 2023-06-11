package services

import (
	"encoding/json"

	"github.com/levigross/grequests"
)

type Joke struct {
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
}

func GetJoke() Joke {
	response, err := grequests.Get("https://v2.jokeapi.dev/joke/Any", nil)
	if err != nil {
		panic(err)
	}
	var data Joke
	err = json.Unmarshal(response.Bytes(), &data)
	if err != nil {
		panic(err)
	}
	return data
}
