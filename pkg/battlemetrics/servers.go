package battlemetrics

import (
	"encoding/json"
	"fmt"

	"github.com/hink/valhallabot/pkg/battlemetrics/internal/pkg/endpoints"
	"github.com/hink/valhallabot/pkg/battlemetrics/pkg/models"
)

func (c *Client) Servers() ([]*models.Server, error) {
	data, err := get(c, endpoints.Servers)
	if err != nil {
		return nil, err
	}

	parsedData := &struct {
		Data []*models.Server `json:"data"`
	}{}
	if err = json.Unmarshal(data, parsedData); err != nil {
		return nil, err
	}

	return parsedData.Data, nil
}

func (c *Client) Server(id int) (*models.Server, error) {
	data, err := get(c, fmt.Sprintf("%s/%d", endpoints.Servers, id))
	if err != nil {
		return nil, err
	}

	parsedData := &struct {
		Data *models.Server `json:"data"`
	}{}
	if err = json.Unmarshal(data, parsedData); err != nil {
		return nil, err
	}

	return parsedData.Data, nil
}
