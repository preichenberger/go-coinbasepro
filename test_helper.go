package coinbase

import (
  "os"
)

func NewTestClient() *Client {
  /*
  secret := os.Getenv("TEST_COINBASE_SECRET") 
  key := os.Getenv("TEST_COINBASE_KEY") 
  passphrase := os.Getenv("TEST_COINBASE_PASSPHRASE") 

  return &Client{
    BaseURL: "https://api-public.sandbox.exchange.coinbase.com",
    Secret: secret,
    Key: key,
    Passphrase: passphrase,
  }*/

  secret := os.Getenv("COINBASE_SECRET") 
  key := os.Getenv("COINBASE_KEY") 
  passphrase := os.Getenv("COINBASE_PASSPHRASE") 

  return NewClient(secret, key, passphrase)
}
