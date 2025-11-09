package main

import (
	"context"
	"log"
	"sync"

	"golang.org/x/time/rate"
)

type APIConnection struct {
	rateLimiter *rate.Limiter
}

func Open() *APIConnection {
	return &APIConnection{
		rateLimiter: rate.NewLimiter(rate.Limit(1), 1),
	}
}

func (a *APIConnection) ReadFile(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	return nil
}

func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	return nil
}

func main() {
	defer log.Println("Done.")

	apiConnection := Open()
	var wg sync.WaitGroup
	wg.Add(20)

	for range 10 {
		go func() {
			defer wg.Done()
			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				log.Printf("can't ReadFile: %v", err)
			}
			log.Println("ReadFile")
		}()
	}

	for range 10 {
		go func() {
			defer wg.Done()
			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("can't ResolveAddress: %v", err)
			}
			log.Println("ResolveAddress")
		}()
	}

	wg.Wait()
}
