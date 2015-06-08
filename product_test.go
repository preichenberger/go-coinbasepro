package coinbase

import(
  "errors"
  "testing"
  "time"
)

func TestGetProducts(t *testing.T) {
  client := NewTestClient()
  products, err := client.GetProducts()
  if err != nil {
    t.Error(err)
  }

  for _, p := range products {
    if StructHasZeroValues(p) {
      t.Error(errors.New("Zero value"))
    } 
  }
}

func TestGetBook(t *testing.T) {
  client := NewTestClient()
  _, err := client.GetBook("BTC-USD", 1)
  if err != nil {
    t.Error(err)
  }
  _, err = client.GetBook("BTC-USD", 2)
  if err != nil {
    t.Error(err)
  }
  _, err = client.GetBook("BTC-USD", 3)
  if err != nil {
    t.Error(err)
  }
}

func TestGetTicker(t *testing.T) {
  client := NewTestClient()
  ticker, err := client.GetTicker("BTC-USD")
  if err != nil {
    t.Error(err)
  }

  if StructHasZeroValues(ticker) {
    t.Error(errors.New("Zero value"))
  } 
}

func TestListTrades(t *testing.T) {
  var trades []Trade
  client := NewTestClient()
  cursor := client.ListTrades("BTC-USD")

  for cursor.HasMore {
    if err := cursor.NextPage(&trades); err != nil {
      t.Error(err)
    } 
    
    for _, a := range trades {  
      if StructHasZeroValues(a) {
        t.Error(errors.New("Zero value"))
      }
    } 
  }
}

func TestGetHistoricRates(t *testing.T) {
  // Test server is busted
  return

  client := NewTestClient()
  params := GetHistoricRatesParams{
    Start: time.Now().Add(-24 * 4  * time.Hour),
    End: time.Now().Add(-24 * 2 * time.Hour),
    Granularity: 1000,
  }

  _, err := client.GetHistoricRates("BTC-USD", params)
  if err != nil {
    t.Error(err)
  }
}

func TestGetStats(t *testing.T) {
  // Test server is busted
  return

  client := NewTestClient()
  stats, err := client.GetStats("BTC-USD")
  if err != nil {
    t.Error(err)
  }
  if StructHasZeroValues(stats) {
    t.Error(errors.New("Zero value"))
  }
}
