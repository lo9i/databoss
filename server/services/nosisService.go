package services

import (
	"encoding/json"
	"fmt"
	"github.com/lo9i/databoss/server/domain"
	"net/http"
)

func NewNosisService(config domain.NosisApiConfiguration) domain.NosisService {
	return &NosisServiceImpl{Config: config}
}

type NosisServiceImpl struct {
	Config     domain.NosisApiConfiguration
	httpClient *http.Client
}

func (c *NosisServiceImpl) Get(id uint64) (*domain.NosisResponse, error) {
	// "https://ws01.nosis.com/api/variables?usuario=62450&token=179281&documento=20256122325&vr=1&Format=json"
	req, err := c.newRequest("GET", fmt.Sprintf("/variables?usuario=%s&token%s&documento=%d&vr=1&format=json", c.Config.User, c.Config.Password, id))
	if err != nil {
		return nil, err
	}
	var user domain.NosisResponse
	_, err = c.do(req, &user)
	return &user, err
}

func (c *NosisServiceImpl) newRequest(method, path string) (*http.Request, error) {
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 Chrome/79.0.3945.88 Safari/537.36")
	return req, nil
}

func (c *NosisServiceImpl) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
