package coinbase

import(
  "errors"
  "testing"
  "time"
)

func TestGetProducts(t *testing.T) {
  client := NewTestClient()
  _, err := client.GetProducts()
  if err != nil {
    t.Error(err)
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
  _, err := client.GetTicker("BTC-USD")
  if err != nil {
    t.Error(err)
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
  }
}

func TestGetHistoricRates(t *testing.T) {
  client := NewTestClient()
  params := GetHistoricRatesParams{
    Start: time.Now().Add(-24 * 4  * time.Hour),
    End: time.Now().Add(-24 * 2 * time.Hour),
    Granularity: 1000,
  }

  historicalRates, err := client.GetHistoricRates("BTC-USD", params)
  if err != nil {
    t.Error(err)
  }

  if len(historicalRates) == 0 {
    t.Error(errors.New("Incorrect size of historical rates"))
  }
}

func TestGetStats(t *testing.T) {
  client := NewTestClient()
  _, err := client.GetStats("BTC-USD")
  if err != nil {
    t.Error(err)
  }
}
