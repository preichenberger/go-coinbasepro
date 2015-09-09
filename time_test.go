package coinbase

import (
	"encoding/json"
	"errors"
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	client := NewTestClient()
	serverTime, err := client.GetTime()
	if err != nil {
		t.Error(err)
	}

	if StructHasZeroValues(serverTime) {
		t.Error(errors.New("Zero value"))
	}
}

func TestTimeUnmarshalJSON(t *testing.T) {
	c := Time{}
	now := time.Now()

	jsonData, err := json.Marshal(now.Format("2006-01-02 15:04:05+00"))
	if err != nil {
		t.Error(err)
	}

	if err = c.UnmarshalJSON(jsonData); err != nil {
		t.Error(err)
	}

	if now.Equal(c.Time()) {
		t.Error(errors.New("Unmarshaled time does not equal original time"))
	}
}
