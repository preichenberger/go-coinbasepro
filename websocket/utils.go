package websocket

import (
	"fmt"
	"strconv"
)

// FormatPrice format price to two decimal places
func FormatPrice(price float64) float64 {
	formattedPrice, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", price), 64)
	return formattedPrice
}
