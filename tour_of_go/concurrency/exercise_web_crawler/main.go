package main

import (
	"fmt"
	"sync"

	"example.com/internal/fetch"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type SafeCache struct {
	mu  sync.Mutex
	val map[string]int
}

var safeCache = SafeCache{
	val: make(map[string]int),
}

func (sc *SafeCache) Cache(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.val[key]++
}

func (sc *SafeCache) IsCached(key string) bool {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	_, ok := sc.val[key]
	return ok
}

func Crawl(url string, depth int, fetcher Fetcher, ch chan int) {
	if depth <= 0 {
		ch <- 1
		return
	}

	if safeCache.IsCached(url) {
		ch <- 1
		return
	} else {
		safeCache.Cache(url)
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		ch <- 1
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	childCh := make(chan int, len(urls))
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
	go Crawl("https://golang.org/", 4, fetch.FakeFetcher, ch)
	<-ch
	fmt.Println("Done!")
}
