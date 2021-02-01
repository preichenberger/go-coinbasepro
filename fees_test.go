package coinbasepro

import (
	"errors"
	"testing"
)

func TestGetFees(t *testing.T) {
	client := NewTestClient()
	fees, err := client.GetFees()
	if err != nil {
		t.Error(err)
	}

	if StructHasZeroValues(fees) {
		t.Error(errors.New("Zero value"))
	}
}
