package gdax

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type Product struct {
	Id             string `json:"id"`
	BaseCurrency   string `json:"base_currency"`
	QuoteCurrency  string `json:"quote_currency"`
	BaseMinSize    string `json:"base_min_size"`
	BaseMaxSize    string `json:"base_max_size"`
	QuoteIncrement string `json:"quote_increment"`
}

type Ticker struct {
	TradeId int    `json:"trade_id,number"`
	Price   string `json:"price"`
	Size    string `json:"size"`
	Time    Time   `json:"time,string"`
	Bid     string `json:"bid"`
	Ask     string `json:"ask"`
	Volume  string `json:"volume"`
}

type Trade struct {
	TradeId int    `json:"trade_id,number"`
	Price   string `json:"price"`
	Size    string `json:"size"`
	Time    Time   `json:"time,string"`
	Side    string `json:"side"`
}

type HistoricRate struct {
	Time   time.Time
	Low    string
	High   string
	Open   string
	Close  string
	Volume string
}

type Stats struct {
	Low          string `json:"low"`
	High         string `json:"high"`
	Open         string `json:"open"`
	Volume       string `json:"volume"`
	Last         string `json:"last"`
	Volume_30Day string `json:"volume_30day"`
}

type BookEntry struct {
	Price          string
	Size           string
	NumberOfOrders int
	OrderId        string
}

type Book struct {
	Sequence int         `json:"sequence"`
	Bids     []BookEntry `json:"bids"`
	Asks     []BookEntry `json:"asks"`
}

type ListTradesParams struct {
	Pagination PaginationParams
}

type GetHistoricRatesParams struct {
	Start       time.Time
	End         time.Time
	Granularity int
}

func (e *BookEntry) UnmarshalJSON(data []byte) error {
	var entry []interface{}

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

	*e = BookEntry{
		Price: priceString,
		Size:  sizeString,
	}

	var stringOrderId string
	numberOfOrdersInt, ok := entry[2].(int)
	if !ok {
		// Try to see if it's a string
		stringOrderId, ok = entry[2].(string)
		if !ok {
			return errors.New("Could not parse 3rd column, tried int and string")
		}
		e.OrderId = stringOrderId

	} else {
		e.NumberOfOrders = numberOfOrdersInt
	}

	return nil
}

func (e *HistoricRate) UnmarshalJSON(data []byte) error {
	var entry []interface{}

	if err := json.Unmarshal(data, &entry); err != nil {
		return err
	}

	t, ok := entry[0].(string)
	if !ok {
		return errors.New("Expected string")
	}

	low, ok := entry[1].(string)
	if !ok {
		return errors.New("Expected string")
	}

	high, ok := entry[2].(string)
	if !ok {
		return errors.New("Expected string")
	}

	open, ok := entry[3].(string)
	if !ok {
		return errors.New("Expected string")
	}

	close, ok := entry[4].(string)
	if !ok {
		return errors.New("Expected string")
	}

	volume, ok := entry[5].(string)
	if !ok {
		return errors.New("Expected string")
	}

	tInt, err := strconv.Atoi(t)
	if err != nil {
		return errors.New("Could not convert epoch string to int")
	}

	*e = HistoricRate{
		Time:   time.Unix(int64(tInt), 0),
		Low:    low,
		High:   high,
		Open:   open,
		Close:  close,
		Volume: volume,
	}

	return nil
}

func (c *Client) GetBook(product string, level int) (Book, error) {
	var book Book

	requestURL := fmt.Sprintf("/products/%s/book?level=%d", product, level)
	_, err := c.Request("GET", requestURL, nil, &book)
	return book, err
}

func (c *Client) GetTicker(product string) (Ticker, error) {
	var ticker Ticker

	requestURL := fmt.Sprintf("/products/%s/ticker", product)
	_, err := c.Request("GET", requestURL, nil, &ticker)
	return ticker, err
}

func (c *Client) ListTrades(product string,
	p ...ListTradesParams) *Cursor {
	paginationParams := PaginationParams{}
	if len(p) > 0 {
		paginationParams = p[0].Pagination
	}

	return NewCursor(c, "GET", fmt.Sprintf("/products/%s/trades", product),
		&paginationParams)
}

func (c *Client) GetProducts() ([]Product, error) {
	var products []Product

	requestURL := fmt.Sprintf("/products")
	_, err := c.Request("GET", requestURL, nil, &products)
	return products, err
}

func (c *Client) GetHistoricRates(product string,
	p ...GetHistoricRatesParams) ([]HistoricRate, error) {
	var historicRates []HistoricRate
	requestURL := fmt.Sprintf("/products/%s/candles", product)
	params := GetHistoricRatesParams{}
	if len(p) > 0 {
		params = p[0]
	}

	if !params.Start.IsZero() && !params.End.IsZero() && params.Granularity != 0 {
		values := url.Values{}
		layout := "2006-01-02T15:04:05Z"
		values.Add("start", params.Start.UTC().Format(layout))
		values.Add("end", params.End.UTC().Format(layout))
		values.Add("granularity", strconv.Itoa(params.Granularity))

		requestURL = fmt.Sprintf("%s?%s", requestURL, values.Encode())
	}

	_, err := c.Request("GET", requestURL, nil, &historicRates)
	return historicRates, err
}

func (c *Client) GetStats(product string) (Stats, error) {
	var stats Stats
	requestURL := fmt.Sprintf("/products/%s/stats", product)
	_, err := c.Request("GET", requestURL, nil, &stats)
	return stats, err
}
