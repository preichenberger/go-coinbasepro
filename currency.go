package coinbasepro

import (
	"fmt"
)

type Details struct {
	Type      string `json:"type"`
	Symbol      string `json:"symbol"`
	NetworkConfirmations      int `json:"network_confirmations"`
	SortOrder int `json:"sort_order"`
	CryptoAddressLink      string `json:"crypto_address_link"`
	CryptoTransactionLink      string `json:"crypto_transaction_link"`
	PushPaymentMethods []string `json:"push_payment_methods"`
	GroupTypes []string `json:"group_types"`
	DisplayName string `json:"display_name"`
	ProcessingTimeSeconds int `json:"processing_time_seconds"`
}

type Currency struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	MinSize string `json:"min_size"`
	Status string `json:"status"`
	Message string `json:"message"`
	MaxPrecision string `json:"max_precision"`
	ConvertableTo []string `json:"convertable_to"`
	Details Details `json:"details"`
}

func (c *Client) GetCurrencies() ([]Currency, error) {
	var currencies []Currency

	url := fmt.Sprintf("/currencies")
	_, err := c.Request("GET", url, nil, &currencies)
	return currencies, err
}
