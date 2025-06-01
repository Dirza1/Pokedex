package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Main() {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		fmt.Printf("There was an issue with the GET request%s", err)
	}
	body, err := io.ReadAll((res.Body))
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("response failed with the status code: %d and \nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	maps := Maps{}
	err = json.Unmarshal(body, &maps)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d,\n %s,\n %v,\n", maps.Count, maps.Next, maps.Previous)
	for _, location := range maps.Results {
		fmt.Printf("%v\n", location.Name)
	}

}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Maps struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}
