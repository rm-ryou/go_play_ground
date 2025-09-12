package main

import (
	"fmt"
	"sync"

	"example.com/internal/fetch"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type SafeCounter struct {
	mu  sync.Mutex
	val map[string]int
}

var safeCounter = &SafeCounter{
	val: make(map[string]int),
}

func (sc *SafeCounter) Cache(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.val[key]++
}

func (sc *SafeCounter) IsCached(key string) bool {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	_, ok := sc.val[key]
	return ok
}

func Crawl(url string, depth int, fetcher Fetcher, ch chan int) {
	if !safeCounter.IsCached(url) {
		safeCounter.Cache(url)
	} else {
		ch <- 1
		return
	}

	if depth <= 0 {
		ch <- 1
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		ch <- 1
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	childCh := make(chan int)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, childCh)
	}

	for range len(urls) {
		<-childCh
	}

	ch <- 1
	return
}

func main() {
	ch := make(chan int)
	go Crawl("https://golang.org/", 4, fetch.DummyData, ch)
	<-ch
	fmt.Println("Done")
}
