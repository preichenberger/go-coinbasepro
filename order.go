package coinbase

import(
  "fmt"
)

type Order struct {
  ClientOID string `json:"client_oid,omitempty"`
  Id string `json:"id"`
  Size float64 `json:"size,string"`
  Funds float64 `json:"funds,string"`
  Price float64 `json:"price,string"`
  Status string `json:"status"`
  Settled bool `json:"settled"`
  Side string `json:"side"`
  ProductId string `json:"product_id"`
  DoneReason string `json:"done_reason"`
  Type string `json:"type"`
  CreatedAt Time `json:"created_at,string"`
}

type ListOrdersParams struct {
  Status string
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


func (c *Client) ListOrders(p ...ListOrdersParams) (*Cursor) {
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
