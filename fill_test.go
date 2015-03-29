package coinbase

import(
  "testing"
)

func TestListFills(t *testing.T) {
  client := NewTestClient()
  cursor := client.ListFills()
  var fills []Fill

  for cursor.HasMore {
    if err := cursor.NextPage(&fills); err != nil {
      t.Error(err)
    }

    for _, _ = range fills {
    }
  }
  params := ListFillsParams{
    ProductId: "BTC-USD",
  }
  cursor = client.ListFills(params)
  for cursor.HasMore {
    if err := cursor.NextPage(&fills); err != nil {
      t.Error(err)
    }

    for _, _ = range fills {
    }
  }

}
