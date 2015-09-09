package coinbase

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Client struct {
	BaseURL    string
	Secret     string
	Key        string
	Passphrase string
}

func NewClient(secret, key, passphrase string) *Client {
	client := Client{
		BaseURL:    "https://api.exchange.coinbase.com",
		Secret:     secret,
		Key:        key,
		Passphrase: passphrase,
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
	req.Header.Add("CB-ACCESS-KEY", c.Key)
	req.Header.Add("CB-ACCESS-PASSPHRASE", c.Passphrase)
	req.Header.Add("CB-ACCESS-TIMESTAMP", timestamp)

	message := fmt.Sprintf("%s%s%s%s", timestamp, method, url,
		string(data))

	sig, err := c.generateSig(message, c.Secret)
	if err != nil {
		return res, err
	}
	req.Header.Add("CB-ACCESS-SIGN", sig)

	client := http.Client{}
	res, err = client.Do(req)
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

func (c *Client) generateSig(message, secret string) (string, error) {
	key, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	signature := hmac.New(sha256.New, key)
	_, err = signature.Write([]byte(message))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature.Sum(nil)), nil
}
