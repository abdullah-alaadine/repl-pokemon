package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}
	cachedData, ok := c.cache.Get(fullURL)
	if ok {
		locationAreasResponse := LocationAreasResponse{}
		err := json.Unmarshal(cachedData, &locationAreasResponse)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return locationAreasResponse, nil
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	locationAreasResponse := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationAreasResponse)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreasResponse, nil
}
