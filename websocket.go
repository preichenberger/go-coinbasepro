package coinbase

import (
	ws "github.com/gorilla/websocket"
)

type MessageChannel chan Message

type WsClient struct {
	WsURL   string
	Channel MessageChannel

	dialer     ws.Dialer
	connection *ws.Conn
}

type WsSubscribeMessage struct {
	Type       string   `json:"type"`
	ProductIds []string `json:"product_ids"`
}

// Creates a WebSocket Client to receive GDAX feed
func NewWsClient(c MessageChannel) *WsClient {
	client := WsClient{
		WsURL:   "wss://ws-feed.gdax.com",
		Channel: c,
	}

	return &client
}

// Creates a WebSocket Client to receive GDAX feed from the Sandbox environment
func NewSandBoxWsClient(c MessageChannel) *WsClient {
	client := NewWsClient(c)
	client.WsURL = "wss://ws-feed-public.sandbox.gdax.com"

	return client
}

// Connects the WSClient to the Websocket endpoint. Returns nil if everything
// went well
func (client *WsClient) Connect() error {
	conn, _, err := client.dialer.Dial(client.WsURL, nil)

	if err != nil {
		return err
	}

	client.connection = conn
	return nil
}

func (client *WsClient) Subscribe(currency_pairs []string) error {
	msg := WsSubscribeMessage{Type: "subscribe", ProductIds: currency_pairs}

	return client.connection.WriteJSON(msg)
}

func (client *WsClient) Run() error {
	var msg Message

	for {
		if err := client.connection.ReadJSON(&msg); err != nil {
			return err
		}
		client.Channel <- msg
	}
}

func (client *WsClient) Close() {
	client.connection.Close()
}
