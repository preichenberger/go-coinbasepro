package coinbase

import(
  "testing"
)

func TestGetTime(t *testing.T) {
  client := NewTestClient()
  _, err := client.GetTime()
  if err != nil {
    t.Error(err)
  }
}
