package coinbasepro

import (
	"fmt"
)

type Deposit struct {
	Currency        string `json:"currency"`
	Amount          string `json:"amount"`
	PaymentMethodID string `json:"payment_method_id"` //payment method id can be determined by calling GetPaymentMethods() function
}

func (c *Client) CreateDeposit(newDeposit *Deposit) (Deposit, error) {
	var savedDeposit Deposit

	url := fmt.Sprintf("/deposits/payment-method")
	_, err := c.Request("POST", url, newDeposit, &savedDeposit)
	return savedDeposit, err
}

type PaymentMethod struct {
	Currency string `json:"currency"`
	Type     string `json:"type"`
	ID       string `json:"ID"`
}

func (c *Client) GetPaymentMethods() ([]PaymentMethod, error) {
	var paymentMethods []PaymentMethod

	url := fmt.Sprintf("/payment-methods")
	_, err := c.Request("GET", url, nil, &paymentMethods)

	return paymentMethods, err

}
