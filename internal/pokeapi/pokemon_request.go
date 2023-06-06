package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "location-area"
	fullURL := baseURL + endpoint

	cachedData, ok := c.cache.Get(fullURL)
	if ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(cachedData, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, data)

	return pokemon, nil
}
