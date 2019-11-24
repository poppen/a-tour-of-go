package main

import (
	"fmt"
	"sync"
)

// Fetcher returns the body of URL and
// a slice of URLs found on that page.
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan int) {
	defer close(ch)
	vi := &visit{urls: make(map[string]int)}
	childSpider(url, depth-1, fetcher, vi)

	return
}

func childSpider(url string, depth int, fetch Fetcher, vi *visit) {
	if depth <= 0 {
		return
	}
	if vi.isVisited(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go func(url string) {
			childSpider(url, depth-1, fetcher, vi)
		}(u)
	}
	return
}

func main() {
	ch := make(chan int)
	Crawl("https://golang.org/", 4, fetcher, ch)
	for {
		if _, ok := <-ch; ok {
			break
		}
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

type visit struct {
	urls map[string]int
	mux  sync.Mutex
}

func (v *visit) isVisited(url string) bool {
	defer v.mux.Unlock()

	v.mux.Lock()
	if _, ok := v.urls[url]; ok {
		return true
	}
	v.urls[url]++
	return false
}
