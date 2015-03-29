package coinbase

import(
  "testing"
)

func TestGetAccounts(t *testing.T) {
  client := NewTestClient()
  _, err := client.GetAccounts()
  if err != nil {
    t.Error(err)
  }
}

func TestGetAccount(t *testing.T) {
  client := NewTestClient()
  accounts, err := client.GetAccounts()
  if err != nil {
    t.Error(err)
  } 

  for _, a := range accounts {
    _, err := client.GetAccount(a.Id)   
    if err != nil {
      t.Error(err)
    }
  }
}
func TestListAccountLedger(t *testing.T) {
  var ledger []LedgerEntry
  client := NewTestClient()
  accounts, err := client.GetAccounts()
  if err != nil {
    t.Error(err)
  }

  for _, a := range accounts {
    cursor := client.ListAccountLedger(a.Id)
    for cursor.HasMore {
      if err := cursor.NextPage(&ledger); err != nil {
        t.Error(err)
      }

      for _, _ = range ledger {
      }
    }
  }
}

func TestListHolds(t *testing.T) {
  var holds []Hold
  client := NewTestClient()
  accounts, err := client.GetAccounts()
  if err != nil {
    t.Error(err)
  }

  for _, a := range accounts {
    cursor := client.ListHolds(a.Id)
    for cursor.HasMore {
      if err := cursor.NextPage(&holds); err != nil {
        t.Error(err)
      }

      for _, _ = range holds {
      }
    }
  }
}
