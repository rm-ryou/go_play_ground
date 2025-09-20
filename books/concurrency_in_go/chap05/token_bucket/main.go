package main

import (
	"context"
	"log"
	"os"
	"sync"

	"golang.org/x/time/rate"
)

func Open() *APIConnection {
	return &APIConnection{
		rateLimiter: rate.NewLimiter(rate.Limit(1), 1),
	}
}

type APIConnection struct {
	rateLimiter *rate.Limiter
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
	defer log.Printf("Done.")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

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
			log.Printf("ReadFile")
		}()
	}

	for range 10 {
		go func() {
			defer wg.Done()
			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("can't ResolveAddress: %v", err)
			}
			log.Printf("ResolveAddress")
		}()
	}

	wg.Wait()
}
