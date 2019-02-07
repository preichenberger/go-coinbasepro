package coinbasepro

import (
	"fmt"
)

type Transfer struct {
	Type              string `json:"type"`
	Amount            string `json:"amount"`
	CoinbaseAccountID string `json:"coinbase_account_id,string"`
}

func (c *Client) CreateTransfer(newTransfer *Transfer) (Transfer, error) {
	var savedTransfer Transfer

	url := fmt.Sprintf("/transfers")
	_, err := c.Request("POST", url, newTransfer, &savedTransfer)
	return savedTransfer, err
}
