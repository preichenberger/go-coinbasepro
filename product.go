package coinbase

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type Product struct {
	Id             string  `json:"id"`
	BaseCurency    string  `json:"base_currency"`
	QuoteCurrency  string  `json:"quote_currency"`
	BaseMinSize    float64 `json:"base_min_size,string"`
	BaseMaxSize    float64 `json:"base_max_size,string"`
	QuoteIncrement float64 `json:"quote_increment,string"`
}

type Ticker struct {
	TradeId int     `json:"trade_id,number"`
	Price   float64 `json:"price,string"`
	Size    float64 `json:"size,string"`
	Time    Time    `json:"time,string"`
}

type Trade struct {
	TradeId int     `json:"trade_id,number"`
	Price   float64 `json:"price,string"`
	Size    float64 `json:"size,string"`
	Time    Time    `json:"time,string"`
	Side    string  `json:"side"`
}

type HistoricRate struct {
	Time   time.Time
	Low    float64
	High   float64
	Open   float64
	Close  float64
	Volume float64
}

type Stats struct {
	Low    float64 `json:"low,string"`
	High   float64 `json:"high,string"`
	Open   float64 `json:"open,string"`
	Volume float64 `json:"volume,string"`
}

type BookEntry struct {
	Price          float64
	Size           float64
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
		Size:  size,
	}

	var stringOrderId string
	numberOfOrdersFloat, ok := entry[2].(float64)
	if !ok {
		// Try to see if it's a string
		stringOrderId, ok = entry[2].(string)
		if !ok {
			return errors.New("Could not parse 3rd column, tried float and string")
		}
		e.OrderId = stringOrderId

	} else {
		e.NumberOfOrders = int(numberOfOrdersFloat)
	}

	return nil
}

func (e *HistoricRate) UnmarshalJSON(data []byte) error {
	var entry []interface{}

	if err := json.Unmarshal(data, &entry); err != nil {
		return err
	}

	timeFloat, ok := entry[0].(float64)
	if !ok {
		return errors.New("Expected float")
	}

	lowFloat, ok := entry[1].(float64)
	if !ok {
		return errors.New("Expected float")
	}

	highFloat, ok := entry[2].(float64)
	if !ok {
		return errors.New("Expected float")
	}

	openFloat, ok := entry[3].(float64)
	if !ok {
		return errors.New("Expected float")
	}

	closeFloat, ok := entry[4].(float64)
	if !ok {
		return errors.New("Expected float")
	}

	volumeFloat, ok := entry[5].(float64)
	if !ok {
		return errors.New("Expected float")
	}

	*e = HistoricRate{
		Time:   time.Unix(int64(timeFloat), 0),
		Low:    lowFloat,
		High:   highFloat,
		Open:   openFloat,
		Close:  closeFloat,
		Volume: volumeFloat,
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
