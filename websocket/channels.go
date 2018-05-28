package websocket

import (
	gdax "github.com/banaio/go-gdax"
)

// NewChannels returns all the channels to subscribe to.
func NewChannels(productIds []string) []gdax.MessageChannel {
	return []gdax.MessageChannel{
		gdax.MessageChannel{
			Name:       "level2",
			ProductIds: productIds,
		},
		// heartbeat is out of sync at the moment
		// seems to be either returning a time in
		// the future or the past by several seconds.
		// gdax.MessageChannel{
		// 	Name:       "heartbeat",
		// 	ProductIds: productIds,
		// },
		// gdax.MessageChannel{
		// 	Name:       "user",
		// 	ProductIds: productIds,
		// },
		// gdax.MessageChannel{
		// 	Name:       "matches",
		// 	ProductIds: productIds,
		// },
		// gdax.MessageChannel{
		// 	Name:       "ticker",
		// 	ProductIds: productIds,
		// },
		// undocument channel ticker_1000
		// gdax.MessageChannel{
		// 	Name:       "ticker_1000",
		// 	ProductIds: productIds,
		// },
		// undocumented channel level2_50
		// gdax.MessageChannel{
		// 	Name:       "level2_50",
		// 	ProductIds: productIds,
		// },
		// gdax.MessageChannel{
		// 	Name:       "full",
		// 	ProductIds: productIds,
		// },
	}
}
