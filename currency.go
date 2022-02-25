package coinbasepro

import (
	"fmt"
)

type Currency struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	MinSize string `json:"min_size"`
}

func (c *Client) GetCurrencies() ([]Currency, error) {
	var currencies []Currency

	url := fmt.Sprintf("/currencies")
	_, err := c.Request("GET", url, nil, &currencies)
	return currencies, err
}

func (c *Client) GetCurrency(currencyID string) (Currency, error) {
	var currency Currency

	url := fmt.Sprintf("/currencies/%s", currencyID)
	_, err := c.Request("GET", url, nil, &currency)
	return currency, err
}
