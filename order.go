package coinbase

import(
  "fmt"
)

type Order struct {
  Id string `json:"id"`
  Size float64 `json:"size,string"`
  Price float64 `json:"price,string"`
  Status string `json:"status"`
  Settled bool `json:"settled"`
  Side string `json:"side"`
  ProductId string `json:"product_id"`
  DoneReason string `json:"done_reason"`
  CreatedAt Time `json:"created_at,string"`
}

type ListOrdersParams struct {
  Pagination PaginationParams
}

func (c *Client) ListOrders(p ...ListOrdersParams) (*Cursor) {
  paginationParams := PaginationParams{}
  if len(p) > 0 {
    paginationParams = p[0].Pagination
  }

  return NewCursor(c, "GET", fmt.Sprintf("/orders"),
    &paginationParams)
}
