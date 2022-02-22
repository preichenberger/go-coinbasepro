package coinbasepro

import (
	"fmt"
)

type Transfer struct {
	ID                string                 `json:"id"`
	Type              string                 `json:"type"`
	CreatedAt         Time                   `json:"created_at"`
	CompletedAt       Time                   `json:"completed_at"`
	CanceledAt        Time                   `json:"canceled_at"`
	ProcessedAt       Time                   `json:"processed_at"`
	Amount            string                 `json:"amount"`
	Details           map[string]interface{} `json:"details"`
	UserNonce         string                 `json:"user_nonce"`
	CoinbaseAccountID string                 `json:"coinbase_account_id,string"`
}

func (c *Client) CreateTransfer(newTransfer *Transfer) (Transfer, error) {
	var savedTransfer Transfer

	url := fmt.Sprintf("/transfers")
	_, err := c.Request("POST", url, newTransfer, &savedTransfer)
	return savedTransfer, err
}

func (c *Client) GetTransfer(transferID string) (Transfer, error) {
	var savedTransfer Transfer

	url := fmt.Sprintf("/transfers/%s", transferID)
	_, err := c.Request("GET", url, nil, &savedTransfer)
	return savedTransfer, err
}
