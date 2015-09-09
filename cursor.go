package coinbase

import (
	"fmt"
)

type Cursor struct {
	Client     *Client
	Pagination *PaginationParams
	Method     string
	Params     interface{}
	URL        string
	HasMore    bool
}

func NewCursor(client *Client, method, url string,
	paginationParams *PaginationParams) *Cursor {
	return &Cursor{
		Client:     client,
		Method:     method,
		URL:        url,
		Pagination: paginationParams,
		HasMore:    true,
	}
}

func (c *Cursor) NextPage(i interface{}) error {
	url := c.URL
	if c.Pagination.Encode("next") != "" {
		url = fmt.Sprintf("%s?%s", c.URL, c.Pagination.Encode("next"))
	}

	res, err := c.Client.Request(c.Method, url, c.Params, i)
	if err != nil {
		c.HasMore = false
		return err
	}

	c.Pagination.Before = res.Header.Get("CB-BEFORE")
	c.Pagination.After = res.Header.Get("CB-AFTER")

	if c.Pagination.Done() {
		c.HasMore = false
	}

	return nil
}
