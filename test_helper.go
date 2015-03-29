package coinbase

import (
  "os"
  "reflect"
)

func NewTestClient() *Client {
  secret := os.Getenv("TEST_COINBASE_SECRET") 
  key := os.Getenv("TEST_COINBASE_KEY") 
  passphrase := os.Getenv("TEST_COINBASE_PASSPHRASE") 

  return &Client{
    BaseURL: "https://api-public.sandbox.exchange.coinbase.com",
    Secret: secret,
    Key: key,
    Passphrase: passphrase,
  }

  return NewClient(secret, key, passphrase)
}

func StructHasZeroValues(i interface{}) bool {
  iv := reflect.ValueOf(i)

  //values := make([]interface{}, v.NumField())

  for i := 0; i < iv.NumField(); i++ {
    field := iv.Field(i)
    if reflect.Zero(field.Type()) == field {
      return true
    }
  }   

  return false
}
