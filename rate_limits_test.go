package gdax

import (
	"errors"
	"testing"
)

var (
	errNotPublicEndpoint  = errors.New("Expecting url to be a public endpoint")
	errNotPrivateEndpoint = errors.New("Expecting url to be a private endpoint")
)

func TestIsPublicEndpoint_PublicEndpoints(t *testing.T) {
	t.Parallel()

	urls := []string{
		"/products",
		"/products/BTC-EUR/book",
		"/products/BTC-EUR/book?level=1",
		"/products/BTC-EUR/book?level=2",
		"/products/BTC-EUR/book?level=3",
		"/products/BTC-USD/book",
		"/products/BTC-USD/book?level=1",
		"/products/BTC-USD/book?level=2",
		"/products/BTC-USD/book?level=3",
		"/products/BTC-EUR/ticker",
		"/products/BTC-USD/ticker",
		"/products/BTC-EUR/trades",
		"/products/BTC-USD/trades",
		"/products/BTC-EUR/candles",
		"/products/BTC-USD/candles",
		"/products/BTC-EUR/stats",
		"/products/BTC-USD/stats",
		"/currencies",
		"/currencies/level1",
		"/currencies/level2",
		"/currencies/level2/level3",
		"/time",
		"/time/level1",
		"/time/level2",
		"/time/level2/level3",
	}

	for _, url := range urls {
		private := !IsPublicEndpoint(url)
		if private {
			t.Error(errNotPublicEndpoint, url)
		}
	}
}

func TestIsPublicEndpoint_PrivateEndpoints(t *testing.T) {
	t.Parallel()

	urls := []string{
		"/accounts",
		"/accounts/9ae6d8dd-ee13-44db-b432-8f8e5912e600",
		"/accounts/39dbeec4-f26d-457a-8d70-2828a1b9ba70",
		"/accounts/d866927a-31a6-445b-b27b-60e6bcbedc24",
		"/accounts/2d833846-82b4-4e11-91b3-94db2f15bb01",
		"/accounts/9ae6d8dd-ee13-44db-b432-8f8e5912e600/ledger",
		"/accounts/39dbeec4-f26d-457a-8d70-2828a1b9ba70/ledger",
		"/accounts/39dbeec4-f26d-457a-8d70-2828a1b9ba70/ledger?after=89782",
		"/accounts/d866927a-31a6-445b-b27b-60e6bcbedc24/ledger",
		"/accounts/d866927a-31a6-445b-b27b-60e6bcbedc24/ledger?after=37947807",
		"/accounts/d866927a-31a6-445b-b27b-60e6bcbedc24/ledger?after=87734",
		"/accounts/d866927a-31a6-445b-b27b-60e6bcbedc24/ledger?after=77609",
		"/accounts/d866927a-31a6-445b-b27b-60e6bcbedc24/ledger?after=67084",
		"/accounts/2d833846-82b4-4e11-91b3-94db2f15bb01/ledger",
		"/accounts/2d833846-82b4-4e11-91b3-94db2f15bb01/ledger?after=37959337",
		"/accounts/2d833846-82b4-4e11-91b3-94db2f15bb01/ledger?after=89672",
		"/accounts/2d833846-82b4-4e11-91b3-94db2f15bb01/ledger?after=86982",
		"/accounts/2d833846-82b4-4e11-91b3-94db2f15bb01/ledger?after=67561",
		"/accounts/2d833846-82b4-4e11-91b3-94db2f15bb01/ledger?after=67199",
		"/accounts/9ae6d8dd-ee13-44db-b432-8f8e5912e600/holds",
		"/accounts/39dbeec4-f26d-457a-8d70-2828a1b9ba70/holds",
		"/accounts/d866927a-31a6-445b-b27b-60e6bcbedc24/holds",
		"/accounts/2d833846-82b4-4e11-91b3-94db2f15bb01/holds",
		"/fills",
		"/fills?after=1651566",
		"/fills?after=11564",
		"/fills?after=8462",
		"/fills?after=431",
		"/fills?product_id=BTC-USD",
		"/fills?after=1651566&product_id=BTC-USD",
		"/fills?after=10529&product_id=BTC-USD",
		"/fills?after=8415&product_id=BTC-USD",
		"/orders",
		"/orders/6a81abeb-821f-49f2-9f33-a3bdebca1087",
		"/orders/fb58ea7a-5a4c-4d48-ab73-59cd2e788820",
		"/orders/7c2bf4b4-ff45-4999-b8eb-25a9a09ea282",
		"/orders/7c2bf4b4-ff45-4999-b8eb-25a9a09ea282",
		"/orders?product_id=LTC-EUR&status=open",
		"/orders?product_id=BTC-USD",
	}

	for _, url := range urls {
		public := IsPublicEndpoint(url)
		if public {
			t.Error(errNotPrivateEndpoint, url)
		}
	}
}
