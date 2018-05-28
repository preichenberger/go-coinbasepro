package orderbook

import (
	"sort"
	"sync"

	gdax "github.com/banaio/go-gdax"
)

// Orderbook represents the state of the orderbook.
type Orderbook struct {
	bids  map[float64]float64
	asks  map[float64]float64
	rwMux *sync.RWMutex
}

// UpdateResult is the result of the call to Update.
type UpdateResult struct {
	DidUpdate bool
	Price     float64
	Side      string
}

// New return new Orderbook
func New() *Orderbook {
	return &Orderbook{
		bids:  map[float64]float64{},
		asks:  map[float64]float64{},
		rwMux: &sync.RWMutex{},
	}
}

// HighestBid returns the highest bid price.
func (o *Orderbook) HighestBid() float64 {
	o.rwMux.RLock()
	defer o.rwMux.RUnlock()

	prices := []float64{}
	for price := range o.bids {
		prices = append(prices, price)
	}
	sort.Float64s(prices)

	highestBid := prices[len(prices)-1]
	return highestBid
}

// LowestAsk returns the lowesk ask price.
func (o *Orderbook) LowestAsk() float64 {
	o.rwMux.RLock()
	defer o.rwMux.RUnlock()

	prices := []float64{}
	for price := range o.asks {
		prices = append(prices, price)
	}
	sort.Float64s(prices)

	lowestAsk := prices[0]
	return lowestAsk
}

// L2update processes an l2update message and returns a tuple
// (bidResult UpdateResult, askResult UpdateResult) indicating
// if HighestBid changed, and if didChangeAsk changed.
//
// message := {
//     "type": "l2update",
//     "product_id": "BTC-EUR",
//     "changes": [
//         ["buy", "6500.09", "0.84702376"],
//         ["sell", "6507.00", "1.88933140"],
//         ["sell", "6505.54", "1.12386524"],
//         ["sell", "6504.38", "0"]
//     ]
// }
func (o *Orderbook) L2update(message *gdax.Message) (*UpdateResult, *UpdateResult) {
	highestBidOld := o.HighestBid()
	lowestAskOld := o.LowestAsk()

	o.rwMux.Lock()
	for _, change := range message.Changes {
		if change.Side == "buy" {
			if change.Size == 0 {
				delete(o.bids, change.Price)
			} else {
				o.bids[change.Price] = change.Size
			}
		} else if change.Side == "sell" {
			if change.Size == 0 {
				delete(o.asks, change.Price)
			} else {
				o.asks[change.Price] = change.Size
			}
		} else {
			// do nothing...
		}
	}
	o.rwMux.Unlock()

	highestBidNew := o.HighestBid()
	didChangeBid := highestBidNew != highestBidOld
	bidResult := &UpdateResult{
		DidUpdate: didChangeBid,
		Price:     highestBidNew,
		Side:      "buy",
	}

	lowestAskNew := o.LowestAsk()
	didChangeAsk := lowestAskNew != lowestAskOld
	askResult := &UpdateResult{
		DidUpdate: didChangeAsk,
		Price:     lowestAskNew,
		Side:      "sell",
	}

	return bidResult, askResult
}

// Snapshot parses the initial snapshot message from
// the websocket.
//
// message := {
//     "type": "snapshot",
//     "product_id": "BTC-EUR",
//     "bids": [["6500.11", "0.45054140"]],
//     "asks": [["6500.15", "0.57753524"]]
// }
func (o *Orderbook) Snapshot(message *gdax.Message) {
	o.rwMux.Lock()
	defer o.rwMux.Unlock()

	for _, entry := range message.Bids {
		o.bids[entry.Price] = entry.Size
	}

	for _, entry := range message.Asks {
		o.asks[entry.Price] = entry.Size
	}
}
