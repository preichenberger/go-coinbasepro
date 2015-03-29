package coinbase

import(
  "net/url"
  "strconv"
)

type PaginationParams struct {
  Limit int
  Before int
  After int
}

func (p *PaginationParams) Encode(direction string) string {
  values := url.Values{}

  if p.Limit > 0 {
    values.Add("limit", strconv.Itoa(p.Limit))
  }
  if p.Before > 0 && direction == "prev" {
    values.Add("before", strconv.Itoa(p.Before)) 
  }
  if p.After > 0 && direction == "next"{
    values.Add("after", strconv.Itoa(p.After))
  }

  return values.Encode()
}

func (p *PaginationParams) Done() bool {
  if p.Before == -1 && p.After == -1 {
    return true
  }

  return false
}
