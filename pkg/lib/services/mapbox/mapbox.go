package mapbox

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	defaultUrl = "https://api.mapbox.com/geocoding/v5/.json"
)

type MapBox struct {
	token      string
	httpClient *http.Client
}

func New(token string) *MapBox {
	return &MapBox{
		httpClient: &http.Client{},
		token:      token,
	}
}

func (mp *MapBox) Validate() error {
	if mp.token == "" {
		return errors.New("MapBox token is empty")
	}

	return nil
}

func (mp *MapBox) Assert() error {
	res, err := mp.httpClient.Get(mp.buildUrl())
	if err != nil {
		return err
	}

	if res.StatusCode == 401 {
		return errors.New("token not authorized to perform request")
	}

	return nil
}

func (mp *MapBox) buildUrl() string {
	return fmt.Sprintf("%s?access_token=%s", defaultUrl, mp.token)
}
