package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (RespDeepLocations, error) {
	url := baseURL + "/location-area/" + locationName

	dat, cached := c.cache.Get(url)
	if !cached {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespDeepLocations{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespDeepLocations{}, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespDeepLocations{}, err
		}

		c.cache.Add(url, dat)
	}

	locationsResp := RespDeepLocations{}
	err := json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespDeepLocations{}, err
	}

	return locationsResp, nil
}
