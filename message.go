package gdax

import (
	"encoding/json"
	"errors"
	"strconv"
)

type Message struct {
	Type          string              `json:"type"`
	ProductId     string              `json:"product_id"`
	ProductIds    []string            `json:"product_ids"`
	TradeId       int                 `json:"trade_id,number"`
	OrderId       string              `json:"order_id"`
	Sequence      int64               `json:"sequence,number"`
	MakerOrderId  string              `json:"maker_order_id"`
	TakerOrderId  string              `json:"taker_order_id"`
	Time          Time                `json:"time,string"`
	RemainingSize float64             `json:"remaining_size,string"`
	NewSize       float64             `json:"new_size,string"`
	OldSize       float64             `json:"old_size,string"`
	Size          float64             `json:"size,string"`
	Price         float64             `json:"price,string"`
	Side          string              `json:"side"`
	Reason        string              `json:"reason"`
	OrderType     string              `json:"order_type"`
	Funds         float64             `json:"funds,string"`
	NewFunds      float64             `json:"new_funds,string"`
	OldFunds      float64             `json:"old_funds,string"`
	Message       string              `json:"message"`
	Bids          []BookEntrySnapshot `json:"bids,omitempty"`
	Asks          []BookEntrySnapshot `json:"asks,omitempty"`
	Changes       []BookChange        `json:"changes,omitempty"`
	LastSize      float64             `json:"last_size,string"`
	BestBid       float64             `json:"best_bid,string"`
	BestAsk       float64             `json:"best_ask,string"`
	Channels      []MessageChannel    `json:"channels"`
	UserId        string              `json:"user_id"`
	ProfileId     string              `json:"profile_id"`
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

type BookEntrySnapshot struct {
	Price float64
	Size  float64
}

// UnmarshalJSON Unmarshals the changes bids and asks in the message received:
//
// data := {
//     "type": "snapshot",
//     "product_id": "BTC-EUR",
//     "bids": [["6500.11", "0.45054140"]],
//     "asks": [["6500.15", "0.57753524"]]
// }
func (e *BookEntrySnapshot) UnmarshalJSON(data []byte) error {
	var entry []interface{}

	if err := json.Unmarshal(data, &entry); err != nil {
		return err
	}

	priceString, ok := entry[0].(string)
	if !ok {
		return errors.New("Expected string")
	}

	sizeString, ok := entry[1].(string)
	if !ok {
		return errors.New("Expected string")
	}

	price, err := strconv.ParseFloat(priceString, 32)
	if err != nil {
		return err
	}

	size, err := strconv.ParseFloat(sizeString, 32)
	if err != nil {
		return err
	}

	*e = BookEntrySnapshot{
		Price: price,
		Size:  size,
	}

	return nil
}

type BookChange struct {
	Side  string
	Price float64
	Size  float64
}

// UnmarshalJSON Unmarshals the changes field in the message received:
//
// data := {
//     "type": "l2update",
//     "product_id": "BTC-EUR",
//     "changes": [
//         ["buy", "6500.09", "0.84702376"],
//         ["sell", "6507.00", "1.88933140"],
//         ["sell", "6505.54", "1.12386524"],
//         ["sell", "6504.38", "0"]
//     ]
// }
func (e *BookChange) UnmarshalJSON(data []byte) error {
	var entry []interface{}

	if err := json.Unmarshal(data, &entry); err != nil {
		return err
	}

	side, ok := entry[0].(string)
	if !ok {
		return errors.New("Expected string")
	}

	priceString, ok := entry[1].(string)
	if !ok {
		return errors.New("Expected string")
	}

	sizeString, ok := entry[2].(string)
	if !ok {
		return errors.New("Expected string")
	}

	price, err := strconv.ParseFloat(priceString, 32)
	if err != nil {
		return err
	}

	size, err := strconv.ParseFloat(sizeString, 32)
	if err != nil {
		return err
	}

	*e = BookChange{
		Side:  side,
		Price: price,
		Size:  size,
	}

	return nil
}
