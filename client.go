package gdax

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"golang.org/x/time/rate"
)

type Client struct {
	BaseURL        string
	Secret         string
	Key            string
	Passphrase     string
	PrivateLimiter *rate.Limiter
	PublicLimiter  *rate.Limiter
	IsSandbox      bool
	HttpClient     *http.Client
}

func New(secret, key, passphrase string, isSandbox bool) (*Client, error) {
	return NewClient(secret, key, passphrase, isSandbox), nil
}

func NewClient(secret, key, passphrase string, isSandbox bool) *Client {
	var baseURL string
	if !isSandbox {
		baseURL = "https://api.gdax.com"
	} else {
		baseURL = "https://api-public.sandbox.gdax.com"
	}

	client := Client{
		BaseURL:        baseURL,
		Secret:         secret,
		Key:            key,
		Passphrase:     passphrase,
		PrivateLimiter: NewRateLimiter(true),
		PublicLimiter:  NewRateLimiter(false),
		IsSandbox:      isSandbox,
		HttpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}

	return &client
}

func (c *Client) Request(method string, url string,
	params, result interface{}) (res *http.Response, err error) {
	var data []byte
	body := bytes.NewReader(make([]byte, 0))

	if params != nil {
		data, err = json.Marshal(params)
		if err != nil {
			return res, err
		}

		body = bytes.NewReader(data)
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, url)
	req, err := http.NewRequest(method, fullURL, body)
	if err != nil {
		return res, err
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// XXX: Sandbox time is off right now
	if os.Getenv("TEST_COINBASE_OFFSET") != "" {
		inc, err := strconv.Atoi(os.Getenv("TEST_COINBASE_OFFSET"))
		if err != nil {
			return res, err
		}

		timestamp = strconv.FormatInt(time.Now().Unix()+int64(inc), 10)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Baylatent Bot 2.0")

	h, err := c.Headers(method, url, timestamp, string(data))
	for k, v := range h {
		req.Header.Add(k, v)
	}

	limiter := c.PrivateLimiter
	if IsPublicEndpoint(url) {
		limiter = c.PublicLimiter
	}
	if err := limiter.Wait(context.Background()); err != nil {
		return nil, err
	}

	res, err = c.HttpClient.Do(req)
	if err != nil {
		return res, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		defer res.Body.Close()
		coinbaseError := Error{}
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&coinbaseError); err != nil {
			return res, err
		}

		return res, error(coinbaseError)
	}

	if result != nil {
		decoder := json.NewDecoder(res.Body)
		if err = decoder.Decode(result); err != nil {
			return res, err
		}
	}

	return res, nil
}

// Headers generates a map that can be used as headers to authenticate a request
func (c *Client) Headers(method, url, timestamp, data string) (map[string]string, error) {
	h := make(map[string]string)
	h["CB-ACCESS-KEY"] = c.Key
	h["CB-ACCESS-PASSPHRASE"] = c.Passphrase
	h["CB-ACCESS-TIMESTAMP"] = timestamp

	message := fmt.Sprintf(
		"%s%s%s%s",
		timestamp,
		method,
		url,
		data,
	)

	sig, err := generateSig(message, c.Secret)
	if err != nil {
		return nil, err
	}
	h["CB-ACCESS-SIGN"] = sig
	return h, nil
}
