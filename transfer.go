package coinbase

import (
	"fmt"
)

type Transfer struct {
	Type              string  `json:"type"`
	Amount            float64 `json:"amount,string"`
	CoinbaseAccountId string  `json:"coinbase_account_id,string"`
}

func (c *Client) CreateTransfer(newTransfer *Transfer) (Transfer, error) {
	var savedTransfer Transfer

	url := fmt.Sprintf("/transfers")
	_, err := c.Request("POST", url, newTransfer, &savedTransfer)
	return savedTransfer, err
}
