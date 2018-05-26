package gdax

import (
	"strings"

	"golang.org/x/time/rate"
)

// https://docs.gdax.com/#rate-limits

type rateLimit struct {
	rate   rate.Limit
	bursts int
}

var publicRateLimit = &rateLimit{
	rate:   3,
	bursts: 6,
}

var privateRateLimit = &rateLimit{
	rate:   5,
	bursts: 10,
}

// NewRateLimiter Create a private or public rate limiter
func NewRateLimiter(isPrivate bool) *rate.Limiter {
	if isPrivate {
		return rate.NewLimiter(privateRateLimit.rate, privateRateLimit.bursts)
	}

	return rate.NewLimiter(publicRateLimit.rate, publicRateLimit.bursts)
}

// https://docs.gdax.com/#market-data
var publicEndpoints = map[string]bool{
	"/products":   true,
	"/currencies": true,
	"/time":       true,
}

// IsPublicEndpoint determines if requestURL is a public endpoint. See https://docs.gdax.com/#market-data.
//
// requestURL should be the path not the entire url, e.g.,
// correct: /time
// incorrect: https://api-public.sandbox.gdax.com/time
func IsPublicEndpoint(requestURL string) bool {
	_, ok := publicEndpoints[requestURL]
	if ok {
		return true
	}

	for url := range publicEndpoints {
		public := strings.HasPrefix(requestURL, url)
		if public {
			return true
		}
	}

	return false
}
