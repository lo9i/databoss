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

func (c *NosisServiceImpl) Get(id string) (*domain.NosisResponse, error) {
	// "https://ws01.nosis.com/api/variables?usuario=62450&token=179281&documento=20256122325&vr=1&Format=json"
	resp, err := http.Get(fmt.Sprintf("%s/variables?usuario=%s&token=%s&documento=%s&vr=1&format=json", c.Config.BaseURL, c.Config.User, c.Config.Password, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var nResp domain.NosisResponse
	err = json.NewDecoder(resp.Body).Decode(&nResp)
	if err != nil {
		return nil, err
	}

	if nResp.Contenido.Resultado.Estado != 200 {
		return nil, fmt.Errorf("%v", nResp.Contenido.Resultado.Novedad)
	}

	return &nResp, err
}
