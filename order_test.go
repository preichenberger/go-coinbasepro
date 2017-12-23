package gdax

import (
	"errors"
	"testing"
)

func TestCreateLimitOrders(t *testing.T) {
	client := NewTestClient()

	order := Order{
		Price:     1.00,
		Size:      1.00,
		Side:      "buy",
		ProductId: "BTC-USD",
	}

	savedOrder, err := client.CreateOrder(&order)
	if err != nil {
		t.Error(err)
	}

	if savedOrder.Id == "" {
		t.Error(errors.New("No create id found"))
	}

	props := []string{"Price", "Size", "Side", "ProductId"}
	_, err = CompareProperties(order, savedOrder, props)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateMarketOrders(t *testing.T) {
	client := NewTestClient()

	order := Order{
		Funds:     10.00,
		Size:      2.00,
		Side:      "buy",
		Type:      "market",
		ProductId: "BTC-USD",
	}

	savedOrder, err := client.CreateOrder(&order)
	if err != nil {
		t.Error(err)
	}

	if savedOrder.Id == "" {
		t.Error(errors.New("No create id found"))
	}

	props := []string{"Price", "Size", "Side", "ProductId"}
	_, err = CompareProperties(order, savedOrder, props)
	if err != nil {
		t.Error(err)
	}
}

func TestCancelOrder(t *testing.T) {
	var orders []Order
	client := NewTestClient()
	cursor := client.ListOrders()
	for cursor.HasMore {
		if err := cursor.NextPage(&orders); err != nil {
			t.Error(err)
		}

		for _, o := range orders {
			if err := client.CancelOrder(o.Id); err != nil {
				if err.Error() != "Order already done" {
					t.Error(err)
				}
			}
		}
	}
}

func TestGetOrder(t *testing.T) {
	client := NewTestClient()

	order := Order{
		Price:     1.00,
		Size:      1.00,
		Side:      "buy",
		ProductId: "BTC-USD",
	}

	savedOrder, err := client.CreateOrder(&order)
	if err != nil {
		t.Error(err)
	}

	getOrder, err := client.GetOrder(savedOrder.Id)
	if err != nil {
		t.Error(err)
	}

	if getOrder.Id != savedOrder.Id {
		t.Error(errors.New("Order ids do not match"))
	}
}

func TestListOrders(t *testing.T) {
	client := NewTestClient()
	cursor := client.ListOrders()
	var orders []Order

	for cursor.HasMore {
		if err := cursor.NextPage(&orders); err != nil {
			t.Error(err)
		}

		for _, o := range orders {
			if StructHasZeroValues(o) {
				t.Error(errors.New("Zero value"))
			}
		}
	}

	cursor = client.ListOrders(ListOrdersParams{Status: "open"})
	for cursor.HasMore {
		if err := cursor.NextPage(&orders); err != nil {
			t.Error(err)
		}

		for _, o := range orders {
			if StructHasZeroValues(o) {
				t.Error(errors.New("Zero value"))
			}
		}
	}
}

func TestCancelAll(t *testing.T) {
	client := NewTestClient()

	if cursor := client.ListOrders(); !cursor.HasMore {
		t.Error("There are no orders to cancel")
	}

	if _, err := client.CancelAll(); err != nil {
		t.Error(err)
	}

	if cursor := client.ListOrders(); cursor.HasMore {
		t.Error("Orders still exist after cancelling all")
	}
}