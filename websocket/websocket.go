package websocket

import (
	"errors"
	"sync"
	"time"

	"fmt"

	gdax "github.com/banaio/go-gdax"
	"github.com/banaio/go-gdax/orderbook"

	ws "github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var (
	errHeartbeatTimeout = errors.New("Heartbeat timeout")
)

// WebSocket a websocket connection to GDAX.
type WebSocket struct {
	client         *gdax.Client
	productIds     []string
	isSandbox      bool
	hearbeat       *gdax.Message
	rwHeartbeatMux *sync.RWMutex
	orderbook      *orderbook.Orderbook
	rwBookMux      *sync.RWMutex
	snapshotReady  chan bool
	hearbeatReady  chan bool
	startTime      time.Time
}

// New returns a new WebSocket
func New(client *gdax.Client, productIds []string, isSandbox bool) *WebSocket {
	if len(productIds) == 0 {
		productIds = []string{
			"BTC-EUR",
		}
	}
	// only market with any activity on the sandbox is BTC-USD
	if len(productIds) == 1 && productIds[0] == "BTC-EUR" && isSandbox {
		productIds = []string{
			"BTC-USD",
		}
	}

	f := &WebSocket{
		client:         client,
		productIds:     productIds,
		isSandbox:      isSandbox,
		rwHeartbeatMux: &sync.RWMutex{},
		orderbook:      orderbook.New(),
		rwBookMux:      &sync.RWMutex{},
		snapshotReady:  make(chan bool, 1),
		hearbeatReady:  make(chan bool, 1),
		startTime:      time.Now(),
	}

	return f
}

// Run start subscription.
func (w *WebSocket) Run(secret, key, passphrase string) {
	w.init()
	wsConn := w.subscribe(secret, key, passphrase)
	w.run(wsConn)
}

func (w *WebSocket) init() {
	go func() {
		<-w.snapshotReady
	}()

	go func() {
		<-w.hearbeatReady
		for {
			w.heartbeatChecker()
			time.Sleep(time.Millisecond * 750)
		}
	}()
}

func (w *WebSocket) subscribe(secret, key, passphrase string) *ws.Conn {
	urlStr := "wss://ws-feed.gdax.com"
	if w.isSandbox {
		urlStr = "wss://ws-feed-public.sandbox.gdax.com"
	}
	wsDialer := &ws.Dialer{}
	wsConn, _, err := wsDialer.Dial(urlStr, nil)

	if err != nil {
		log.WithFields(log.Fields{
			"err":    err,
			"urlStr": urlStr,
		}).Fatal("WebSocket:subscribe")
	}

	channels := NewChannels(w.productIds)
	if secret != "" && key != "" && passphrase != "" {
		msg, err := gdax.Message{
			Type:     "subscribe",
			Channels: channels,
		}.Sign(secret, key, passphrase)
		if err != nil {
			log.WithFields(log.Fields{
				"err":      err,
				"msg":      fmt.Sprintf("%+v", msg),
				"channels": fmt.Sprintf("%+v", channels),
			}).Fatal("WebSocket:subscribe")
		}

		if err := wsConn.WriteJSON(msg); err != nil {
			log.WithFields(log.Fields{
				"err":      err,
				"msg":      fmt.Sprintf("%+v", msg),
				"channels": fmt.Sprintf("%+v", channels),
			}).Fatal("WebSocket:subscribe")
		}
	} else {
		msg := &gdax.Message{
			Type:     "subscribe",
			Channels: channels,
		}
		if err := wsConn.WriteJSON(msg); err != nil {
			log.WithFields(log.Fields{
				"err":      err,
				"msg":      fmt.Sprintf("%+v", msg),
				"channels": fmt.Sprintf("%+v", channels),
			}).Fatal("WebSocket:subscribe")
		}
	}
	log.WithFields(log.Fields{
		"urlStr":     urlStr,
		"productIds": w.productIds,
		"channels":   channels,
	}).Info("WebSocket:subscribe")

	return wsConn
}

func (w *WebSocket) run(wsConn *ws.Conn) {
	for {
		message := &gdax.Message{}
		if err := wsConn.ReadJSON(message); err != nil {
			log.WithFields(log.Fields{
				"err":     err,
				"message": fmt.Sprintf("%+v", message),
			}).Error("WebSocket:run")
			continue
		}

		if message.Type == "error" {
			log.WithFields(log.Fields{
				"Reason":  message.Reason,
				"Message": message.Message,
				// "message": fmt.Sprintf("%+v", message),
			}).Fatal("WebSocket:run")
		}

		func() {
			w.process(message)
		}()
	}
}

func (w *WebSocket) process(message *gdax.Message) {
	if message.Type == "heartbeat" {
		w.heartbeat(message)
	} else if message.Type == "snapshot" {
		w.snapshot(message)
	} else if message.Type == "l2update" {
		w.l2update(message)
	} else if message.Type == "match" {
		w.match(message)
	} else if message.Type == "received" {
		w.received(message)
	} else if message.Type == "open" {
		w.open(message)
	} else if message.Type == "done" {
		w.done(message)
	} else {
		log.WithFields(log.Fields{
			"Type": message.Type,
			"Time": message.Time.Time().Format(time.RFC3339Nano),
			// "message": fmt.Sprintf("%+v", message),
		}).Info("WebSocket:process")
	}
}

// Orderbook returns a copy of the Orderbook.
func (w *WebSocket) Orderbook() orderbook.Orderbook {
	w.rwBookMux.RLock()
	defer w.rwBookMux.RUnlock()

	return *w.orderbook
}

func (w *WebSocket) snapshot(message *gdax.Message) {
	log.WithFields(log.Fields{
		"Type": message.Type,
		"Time": message.Time.Time().Format(time.RFC3339Nano),
		// "highestBid": highestBid,
		// "lowestAsk":  lowestAsk,
		// "message": fmt.Sprintf("%+v", message),
	}).Info("WebSocket:snapshot")

	w.rwBookMux.Lock()
	w.orderbook.Snapshot(message)

	highestBid := w.orderbook.HighestBid()
	lowestAsk := w.orderbook.LowestAsk()

	w.sendHighestBid(highestBid)
	w.sendLowestAsk(lowestAsk)
	w.rwBookMux.Unlock()

	w.snapshotReady <- true
}

func (w *WebSocket) l2update(message *gdax.Message) {
	// log.WithFields(log.Fields{
	// 	"Type": message.Type,
	// 	"Time": message.Time.Time().Format(time.RFC3339Nano),
	// 	// "message": fmt.Sprintf("%+v", message),
	// }).Info("WebSocket:l2update")

	w.rwBookMux.RLock()
	defer w.rwBookMux.RUnlock()

	bidResult, askResult := w.orderbook.L2update(message)
	if bidResult.DidUpdate {
		w.sendHighestBid(bidResult.Price)
	}
	if askResult.DidUpdate {
		w.sendLowestAsk(askResult.Price)
	}
}

func (w *WebSocket) sendHighestBid(highestBid float64) {
	log.WithFields(log.Fields{
		"HighestBid": FormatPrice(highestBid),
	}).Info("WebSocket:sendHighestBid")
}

func (w *WebSocket) sendLowestAsk(lowestAsk float64) {
	log.WithFields(log.Fields{
		"LowestAsk": FormatPrice(lowestAsk),
	}).Info("WebSocket:sendLowestAsk")
}

func (w *WebSocket) heartbeat(message *gdax.Message) {
	log.WithFields(log.Fields{
		"Type": message.Type,
		"Time": message.Time.Time().Format(time.RFC3339Nano),
		// "message": fmt.Sprintf("%+v", message),
	}).Warn("WebSocket:heartbeat")

	w.rwHeartbeatMux.Lock()
	w.hearbeat = message
	w.rwHeartbeatMux.Unlock()

	select {
	case <-w.hearbeatReady:
		return
	default:
		w.hearbeatReady <- true
	}
}

func (w *WebSocket) heartbeatChecker() {
	w.rwHeartbeatMux.RLock()
	last := w.hearbeat.Time.Time()
	w.rwHeartbeatMux.RUnlock()

	now := time.Now().UTC()
	diff := now.Sub(last)
	if diff > time.Second || diff < time.Second*0 {
		log.WithFields(log.Fields{
			"err":  errHeartbeatTimeout,
			"now":  now.Format(time.RFC3339Nano),
			"last": last.Format(time.RFC3339Nano),
			"diff": diff.Seconds(),
		}).Error("WebSocket:heartbeatChecker")
	} else {
		// log.WithFields(log.Fields{
		// 	"now":  now.Format(time.RFC3339Nano),
		// 	"last": last.Format(time.RFC3339Nano),
		// 	"diff": diff.Seconds(),
		// }).Info("WebSocket:heartbeatChecker")
	}
}

func (w *WebSocket) match(message *gdax.Message) {
	log.WithFields(log.Fields{
		"Type": message.Type,
		"Time": message.Time.Time().Format(time.RFC3339Nano),
		// "message": fmt.Sprintf("%+v", message),
	}).Info("WebSocket:match")
}

func (w *WebSocket) received(message *gdax.Message) {
	log.WithFields(log.Fields{
		"Type": message.Type,
		"Time": message.Time.Time().Format(time.RFC3339Nano),
		// "message": fmt.Sprintf("%+v", message),
	}).Info("WebSocket:received")
}

func (w *WebSocket) open(message *gdax.Message) {
	log.WithFields(log.Fields{
		"Type": message.Type,
		"Time": message.Time.Time().Format(time.RFC3339Nano),
		// "message": fmt.Sprintf("%+v", message),
	}).Info("WebSocket:open")
}

func (w *WebSocket) done(message *gdax.Message) {
	log.WithFields(log.Fields{
		"Type": message.Type,
		"Time": message.Time.Time().Format(time.RFC3339Nano),
		// "message": fmt.Sprintf("%+v", message),
	}).Info("WebSocket:done")
}
