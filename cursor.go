package coinbase

import(
  "fmt"
  "strconv"
)

type Cursor struct {
  Client *Client
  Pagination *PaginationParams
  Method string
  Params interface{}
  URL string
  HasMore bool
}

func NewCursor(client *Client, method, url string,
  paginationParams *PaginationParams) *Cursor {
  return &Cursor{
    Client: client,
    Method: method,
    URL: url,
    Pagination: paginationParams,
    HasMore: true,
  }
}

func (c *Cursor) NextPage(i interface{}) error {
  url := c.URL
  if c.Pagination.Encode("next") != "" {
    url = fmt.Sprintf("%s?%s", c.URL, c.Pagination.Encode("next"))  
  }

  res, err := c.Client.Request(c.Method, url, c.Params, i)
  if err != nil {
    return err
  }

  println(res.Header.Get("CB-BEFORE"))
  println(res.Header.Get("CB-AFTER"))

  if res.Header.Get("CB-BEFORE") == "" { 
    c.Pagination.Before = -1
  } else {
    before, err := strconv.Atoi(res.Header.Get("CB-BEFORE"))
    if err != nil {
      return err
    } 
    c.Pagination.Before = before
  }

  if res.Header.Get("CB-AFTER") == "" { 
    c.Pagination.After = -1
  } else {
    after, err := strconv.Atoi(res.Header.Get("CB-AFTER"))
    if err != nil {
      return err
    }
    c.Pagination.After = after
  }

  if c.Pagination.Done() {
    println("FINISHED")
    c.HasMore = false
  }
  
  return nil
}
