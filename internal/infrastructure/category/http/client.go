package categoryhttpclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	categorydto "github.com/Oleja123/dcaa-property/pkg/dto/category"
	myErrors "github.com/Oleja123/dcaa-property/pkg/errors"
)

type Client struct {
	client *http.Client
	apiUrl string
}

func (c *Client) FindOne(id int) (categorydto.CategoryDTO, error) {
	resp, err := c.client.Get(c.apiUrl + fmt.Sprintf("/%d", id))
	var category categorydto.CategoryDTO
	if err != nil {
		return category, myErrors.ErrInternalError
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return category, fmt.Errorf("категория по id %d отсутствует: %w", id, myErrors.ErrNotFound)
	}

	if err := json.NewDecoder(resp.Body).Decode(&category); err != nil {
		return category, fmt.Errorf("ошибка декодирования категории с id: %d: %w", id, myErrors.ErrInternalError)
	}

	return category, nil
}

func NewClient(apiUrl string) *Client {
	return &Client{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		apiUrl: apiUrl,
	}
}
