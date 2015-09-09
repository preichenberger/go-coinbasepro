package coinbase

import (
	"fmt"
)

type Order struct {
	Type      string  `json:"type"`
	Size      float64 `json:"size,string"`
	Side      string  `json:"side"`
	ProductId string  `json:"product_id"`
	ClientOID string  `json:"client_oid,omitempty"`
	Stp       string  `json:"stp,omitempty"`
	// Limit Order
	Price       float64 `json:"price,string,omitempty"`
	TimeInForce string  `json:"time_in_force,omitempty"`
	PostOnly    bool    `json:"post_only,omitempty"`
	// Market Order
	Funds float64 `json:"funds,string,omitempty"`
	// Response Fields
	Id         string `json:"id"`
	Status     string `json:"status,omitempty"`
	Settled    bool   `json:"settled,omitempty"`
	DoneReason string `json:"done_reason,omitempty"`
	CreatedAt  Time   `json:"created_at,string,omitempty"`
}

type ListOrdersParams struct {
	Status     string
	Pagination PaginationParams
}

func (c *Client) CreateOrder(newOrder *Order) (Order, error) {
	var savedOrder Order

	if len(newOrder.Type) == 0 {
		newOrder.Type = "limit"
	}

	url := fmt.Sprintf("/orders")
	_, err := c.Request("POST", url, newOrder, &savedOrder)
	return savedOrder, err
}

func (c *Client) CancelOrder(id string) error {
	url := fmt.Sprintf("/orders/%s", id)
	_, err := c.Request("DELETE", url, nil, nil)
	return err
}

func (c *Client) GetOrder(id string) (Order, error) {
	var savedOrder Order

	url := fmt.Sprintf("/orders/%s", id)
	_, err := c.Request("GET", url, nil, &savedOrder)
	return savedOrder, err
}

func (c *Client) ListOrders(p ...ListOrdersParams) *Cursor {
	paginationParams := PaginationParams{}
	if len(p) > 0 {
		paginationParams = p[0].Pagination
		if p[0].Status != "" {
			paginationParams.AddExtraParam("status", p[0].Status)
		}
	}

	return NewCursor(c, "GET", fmt.Sprintf("/orders"),
		&paginationParams)
}
