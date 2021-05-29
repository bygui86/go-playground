package main

import (
	"time"

	"github.com/bygui86/go-playground/ratelimiter"
)

func main() {
	r, err := ratelimiter.NewMaxConcurrencyRateLimiter(&ratelimiter.Config{
		Limit:            2,
		TokenResetsAfter: 10 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	ratelimiter.DoWork(r, 15)
}
