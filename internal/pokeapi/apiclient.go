package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{httpClient: http.Client{}}
}

func (c *Client) GetLocations(wanturl *string) (PokeResponse, error) {

	var url string

	if wanturl == nil {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = *wanturl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokeResponse{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeResponse{}, err
	}
	defer res.Body.Close()

	locs := PokeResponse{}

	err = json.Unmarshal(body, &locs)
	if err != nil {
		return PokeResponse{}, err
	}

	return locs, nil
}

func (c *Client) GetPokemon(wanturl *string) (PokeLocation, error) {

	var url string

	if len(*wanturl) == 0 {
		return PokeLocation{}, fmt.Errorf("Did not provide location")
	} else {
		url = *wanturl
	}

	req, err := http.NewRequest("GET", url, nil)
	fmt.Println(req)
	if err != nil {
		return PokeLocation{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokeLocation{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeLocation{}, err
	}
	defer res.Body.Close()

	locs := PokeLocation{}

	err = json.Unmarshal(body, &locs)
	if err != nil {
		return PokeLocation{}, err
	}

	return locs, nil
}
