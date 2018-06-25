package gdax

type Message struct {
	Type          string           `json:"type"`
	ProductId     string           `json:"product_id"`
	ProductIds    []string         `json:"product_ids"`
	TradeId       int              `json:"trade_id,number"`
	OrderId       string           `json:"order_id"`
	Sequence      int64            `json:"sequence,number"`
	MakerOrderId  string           `json:"maker_order_id"`
	TakerOrderId  string           `json:"taker_order_id"`
	Time          Time             `json:"time,string"`
	RemainingSize string           `json:"remaining_size"`
	NewSize       string           `json:"new_size"`
	OldSize       string           `json:"old_size"`
	Size          string           `json:"size"`
	Price         string           `json:"price"`
	Side          string           `json:"side"`
	Reason        string           `json:"reason"`
	OrderType     string           `json:"order_type"`
	Funds         string           `json:"funds"`
	NewFunds      string           `json:"new_funds"`
	OldFunds      string           `json:"old_funds"`
	Message       string           `json:"message"`
	Bids          [][]string       `json:"bids,omitempty"`
	Asks          [][]string       `json:"asks,omitempty"`
	Changes       [][]string       `json:"changes,omitempty"`
	LastSize      string           `json:"last_size"`
	BestBid       string           `json:"best_bid"`
	BestAsk       string           `json:"best_ask"`
	Channels      []MessageChannel `json:"channels"`
	UserId        string           `json:"user_id"`
	ProfileId     string           `json:"profile_id"`
}

type MessageChannel struct {
	Name       string   `json:"name"`
	ProductIds []string `json:"product_ids"`
}

type SignedMessage struct {
	Message
	Key        string `json:"key"`
	Passphrase string `json:"passphrase"`
	Timestamp  string `json:"timestamp"`
	Signature  string `json:"signature"`
}
