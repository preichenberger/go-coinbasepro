package coinbase

type Message struct {
  Type string `json:"type"`
  TradeId int `json:"trade_id,number"`
  OrderId string `json:"order_id"`
  Sequence int `json:"sequence,number"`
  Time Time `json:"time,string"`
  Size float64 `json:"size,string"`
  Price float64 `json:"price,string"`
  Side string `json:"side"`
  Reason string `json:"reason"`
}
