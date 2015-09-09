package coinbase

import (
	"fmt"
)

type Currency struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	MinSize float64 `json:"min_size,string"`
}

func (c *Client) GetCurrencies() ([]Currency, error) {
	var currencies []Currency

	url := fmt.Sprintf("/currencies")
	_, err := c.Request("GET", url, nil, &currencies)
	return currencies, err
}
