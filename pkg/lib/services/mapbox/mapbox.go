package mapbox

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	defaultUrl = "https://api.mapbox.com/geocoding/v5/.json"
)

type mapBox struct {
	token string
	httpClient *http.Client
}

func New(token string) *mapBox {
	return &mapBox{
		httpClient: &http.Client{},
		token:token,
	}
}

func (mp *mapBox) Validate() error {
	if mp.token == "" {
		return errors.New("MapBox token is empty")
	}

	return nil
}

func (mp *mapBox) Assert() error {
	res, err := mp.httpClient.Get(mp.buildUrl())
	if err != nil {
		return err
	}

	if res.StatusCode == 401 {
		return errors.New("token not authorized to perform request")
	}

	return nil
}

func (mp *mapBox) buildUrl() string {
	return fmt.Sprintf("%s?access_token=%s", defaultUrl, mp.token)
}
