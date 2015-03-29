package coinbase

import(
  "errors"
  "testing"
)

func TestGetTime(t *testing.T) {
  client := NewTestClient()
  serverTime, err := client.GetTime()
  if err != nil {
    t.Error(err)
  }

  if StructHasZeroValues(serverTime) {
    t.Error(errors.New("Zero value"))
  }
}
