package foac

import (
	"encoding/json"
	"net/http"
	"strings"
)

const baseURL = "https://foaas.com/"

type client struct {
	http *http.Client
}

type response struct {
	Message  string `json:"message"`
	Subtitle string `json:"subtitle"`
}

func (c *client) makeRequest(params ...string) (*response, error) {
	req, err := http.NewRequest(http.MethodGet, baseURL+strings.Join(params, "/"), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := c.http.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	target := &response{}

	if err = json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}

	return target, nil
}
