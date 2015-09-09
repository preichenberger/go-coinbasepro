package coinbase

type Message struct {
	Type          string  `json:"type"`
	TradeId       int     `json:"trade_id,number"`
	OrderId       string  `json:"order_id"`
	Sequence      int     `json:"sequence,number"`
	MakerOrderId  string  `json:"maker_order_id"`
	TakerOrderId  string  `json:"taker_order_id"`
	Time          Time    `json:"time,string"`
	RemainingSize float64 `json:"remaining_size,string"`
	NewSize       float64 `json:"new_size,string"`
	OldSize       float64 `json:"old_size,string"`
	Size          float64 `json:"size,string"`
	Price         float64 `json:"price,string"`
	Side          string  `json:"side"`
	Reason        string  `json:"reason"`
	OrderType     string  `json:"order_type"`
	Funds         float64 `json:"funds,string"`
	NewFunds      float64 `json:"new_funds,string"`
	OldFunds      float64 `json:"old_funds,string"`
}
