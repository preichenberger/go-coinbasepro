package coinbase

import(
  "testing"
)

func TestGetCurrencies(t *testing.T) {
  client := NewTestClient()
  _, err := client.GetCurrencies()
  if err != nil {
    t.Error(err)
  }
}
