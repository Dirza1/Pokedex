package pokeapi

import (
	"fmt"
	"net/http"
)

func Main() {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		fmt.Printf("There was an issue with the GET request%s", err)
	}
	fmt.Print(res)
}
