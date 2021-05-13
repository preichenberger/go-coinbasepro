package coinbasepro

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type ServerTime struct {
	ISO   string  `json:"iso"`
	Epoch float64 `json:"epoch,number"`
}

func (c *Client) GetTime() (ServerTime, error) {
	var serverTime ServerTime

	url := fmt.Sprintf("/time")
	_, err := c.Request("GET", url, nil, &serverTime)
	return serverTime, err
}

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {
	var err error
	var parsedTime time.Time

	if string(data) == "null" {
		*t = Time(time.Time{})
		return nil
	}

	layouts := []string{
		"2006-01-02 15:04:05+00",
		"2006-01-02T15:04:05-07:00",
		"2006-01-02T15:04:05.999999Z",

		"2006-01-02 15:04:05.999999",
		"2006-01-02T15:04:05Z",
		"2006-01-02 15:04:05.999999+00"}
	for _, layout := range layouts {
		parsedTime, err = time.Parse(layout,
			strings.Replace(string(data), "\"", "", -1))
		if err != nil {
			continue
		}

		break
	}
	if parsedTime.IsZero() {
		return err
	}

	*t = Time(parsedTime)

	return nil
}

// MarshalJSON marshal time back to time.Time for json encoding
func (t Time) MarshalJSON() ([]byte, error) {
	return t.Time().MarshalJSON()
}

func (t *Time) Time() time.Time {
	return time.Time(*t)
}

// Scan implements the sql.Scanner interface for database deserialization.
func (t *Time) Scan(value interface{}) error {
	timeValue, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("failed to deserialize time: %#v", value)
	}
	*t = Time(timeValue)
	return nil
}

// Value implements the driver.Valuer interface for database serialization.
func (t Time) Value() (driver.Value, error) {
	return t.Time(), nil
}
