package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	url := "https://go.dev"
	bid := bidOn(ctx, url)
	fmt.Println("Bidding at", bid.Price)
}

func bestBid(url string) Bid {
	time.Sleep(20 * time.Millisecond)
	return Bid{
		AdURL: "http://adsЯus.com/ad7",
		Price: 7,
	}
}

type Bid struct {
	AdURL string
	Price int
}

var defaultBid = Bid{
	AdURL: "http://adsЯus.com/default",
	Price: 3,
}

func bidOn(ctx context.Context, url string) Bid {
	// first one to finish will have to bid either ctx timeout or when bestBid finishes calculating the best bid
	ch := make(chan Bid)
	go func() {
		ch <- bestBid(url)
	}()
	select {
	case bid := <-ch:
		return bid
	case <-ctx.Done():
		return defaultBid
	}
}
