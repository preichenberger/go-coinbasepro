Go Coinbase Exchange [![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/preichenberger/go-coinbase-exchange) [![Build Status](https://travis-ci.org/preichenberger/go-coinbase-exchange.svg?branch=master)](https://travis-ci.org/preichenberger/go-coinbase-exchange)
========

## Summary

Go client for [Coinbase Exchange](https://exchange.coinbase.com)

## Installation

```sh
go get github.com/preichenberger/go-coinbase-exchange
```

## Documentation


### Setup
How to create a client:

```go

import (
  "os"
  exchange "github.com/preichenberger/go-coinbase-exchange"
)

secret := os.Getenv("COINBASE_SECRET") 
key := os.Getenv("COINBASE_KEY") 
passphrase := os.Getenv("COINBASE_PASSPHRASE") 

// or unsafe hardcode way
secret = "exposedsecret"
key = "exposedkey"
passphrase = "exposedpassphrase"

client := NewClient(secret, key, passphrase)
```

### Cursor
This library uses a cursor pattern so you don't have to keep track of pagination.

```go
var orders []exchange.Order
cursor = client.ListOrders()

for cursor.HasMore {
  if err := cursor.NextPage(&orders); err != nil {
    println(err.Error())
    return
  }

  for _, o := range orders {
    println(o.Id) 
  }
}

```

### Websockets
Listen for websocket messages

```go
  import(
    ws "github.com/gorilla/websocket"
  )

  var wsDialer ws.Dialer
  wsConn, _, err := wsDialer.Dial("wss://ws-feed.exchange.coinbase.com", nil)
  if err != nil {
    println(err.Error())
  }

  subscribe := map[string]string{
    "type": "subscribe",
    "product_id": "BTC-USD",
  }
  if err := wsConn.WriteJSON(subscribe); err != nil {
    println(err.Error())
  }

  message:= exchange.Message{}
  for true {
    if err := wsConn.ReadJSON(&message); err != nil {
      println(err.Error())
      break
    }

    if message.Type == "match" {
      println("Got a match")
    }
  }

```

### Examples
This library supports all public and private endpoints

For full details on functionality, see [GoDoc](http://godoc.org/github.com/preichenberger/go-coinbase-exchange) documentation.

Get Accounts:
```go
  accounts, err := client.GetAccounts()
  if err != nil {
    println(err.Error())
  }

  for _, a := range accounts {
    println(a.Balance)
  }
```

List Account Ledger:
```go
  var ledger []exchange.LedgerEntry

  accounts, err := client.GetAccounts()
  if err != nil {
    println(err.Error())
  }

  for _, a := range accounts {
    cursor := client.ListAccountLedger(a.Id)
    for cursor.HasMore {
      if err := cursor.NextPage(&ledger); err != nil {
        println(err.Error())
      }

      for _, e := range ledger {
        println(e.Amount)
      }
  }
```

Create an Order:
```go
  order := exchange.Order{
    Price: 1.00,   
    Size: 1.00,   
    Side: "buy",   
    ProductId: "BTC-USD",   
  }

  savedOrder, err := client.CreateOrder(&order)
  if err != nil {
    println(err.Error())
  }

  println(savedOrder.Id)
```

Transfer funds:
```go
  transfer := exchange.Transfer {
    Type: "deposit",
    Amount: 1.00,   
  }

  savedTransfer, err := client.CreateTransfer(&transfer)
  if err != nil {
    println(err.Error())
  }
```

Get Trade history:
```go
  var trades []exchange.Trade
  cursor := client.ListTrades("BTC-USD")

  for cursor.HasMore {
    if err := cursor.NextPage(&trades); err != nil {
      for _, t := range trades {
        println(trade.CoinbaseId)
      }
    } 
  }
```

### Testing
To test with Coinbase's public sandbox set the following environment variables:
  - TEST_COINBASE_SECRET
  - TEST_COINBASE_KEY
  - TEST_COINBASE_PASSPHRASE

Then run `go test`
```sh
TEST_COINBASE_SECRET=secret TEST_COINBASE_KEY=key TEST_COINBASE_PASSPHRASE=passphrase go test
```
