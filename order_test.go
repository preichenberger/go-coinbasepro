package coinbase

import(
  "testing"
)

func TestListOrders(t *testing.T) {
  client := NewTestClient()
  cursor := client.ListOrders()
  var orders []Order
  
  for cursor.HasMore {
    if err := cursor.NextPage(&orders); err != nil {
      t.Error(err)
    }

    for _, _ = range orders {
    }
  }
}
