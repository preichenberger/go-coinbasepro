package gdax

import (
	"fmt"
)

type WithdrawalCrypto struct {
	Currency      string `json:"currency"`
	Amount        string `json:"amount"`
	CryptoAddress string `json:"crypto_address"`
}

type WithdrawalCoinbase struct {
	Currency          string `json:"currency"`
	Amount            string `json:"amount"`
	CoinbaseAccountId string `json:"coinbase_account_id"`
}

func (c *Client) CreateWithdrawalCrypto(newWithdrawCrypto *WithdrawalCrypto) (WithdrawalCrypto, error) {
	var savedWithdrawal WithdrawalCrypto
	url := fmt.Sprintf("/withdrawals/crypto")
	_, err := c.Request("POST", url, newWithdrawCrypto, &savedWithdrawal)
	return savedWithdrawal, err
}

func (c *Client) CreateWithdrawalCoinbase(newWithdrawCoinbase *WithdrawalCoinbase) (WithdrawalCoinbase, error) {
	var savedWithdrawal WithdrawalCoinbase
	url := fmt.Sprintf("/withdrawals/coinbase-account")
	_, err := c.Request("POST", url, newWithdrawCoinbase, &savedWithdrawal)
	return savedWithdrawal, err
}
