package coinbase

import(
  "fmt"
)

type Fill struct {
  TradeId int `json:"trade_id,int"`
  ProductId string `json:"product_id"`
  Price float64 `json:"price,string"`
  Size float64 `json:"size,string"`
  FillId string `json:"order_id"`
  CreatedAt Time `json:"created_at,string"`
  Fee float64 `json:"fee,string"`
  Settled bool `json:"settled"`
  Side string `json:"side"`
  Liquidity string `json:"liquidity"`
}

type ListFillsParams struct {
  OrderId string
  ProductId string
  Pagination PaginationParams
}

func (c *Client) ListFills(p ...ListFillsParams) (*Cursor) {
  paginationParams := PaginationParams{}
  if len(p) > 0 {
    paginationParams = p[0].Pagination
    if p[0].OrderId != "" {
      paginationParams.AddExtraParam("order_id", p[0].OrderId)
    }
    if p[0].ProductId != "" {
      paginationParams.AddExtraParam("product_id", p[0].ProductId)
    }
  }


  return NewCursor(c, "GET", fmt.Sprintf("/fills"),
    &paginationParams)
}
