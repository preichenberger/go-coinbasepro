package coinbase

import(
  "encoding/json"
  "errors"
  "strconv"
)

type BookEntry struct {
  Price float64
  Size float64
  Orders int
}

type Book struct {
  Sequence int `json:"sequence"`
  Bids []BookEntry `json:"bids"`
  Asks []BookEntry `json:"asks"`
}

func(e *BookEntry) UnmarshalJSON(data []byte) error {
  var entry[]interface{}

  if err := json.Unmarshal(data, &entry); err != nil {
    return err
  }

  priceString, ok := entry[0].(string)
  if !ok {
    return errors.New("Expected string")
  }
  
  sizeString, ok := entry[1].(string)
  if !ok {
    return errors.New("Expected string")
  }
  
  ordersFloat, ok := entry[2].(float64)
  if !ok {
    return errors.New("Expected float64")
  }

  price, err := strconv.ParseFloat(priceString, 32)
  if err != nil {
    return err
  }
  
  size, err := strconv.ParseFloat(sizeString, 32)
  if err != nil {
    return err
  }
  
  *e = BookEntry{
    Price: price,
    Size: size,
    Orders: int(ordersFloat),
  }

  return nil
}
