package main

import (
	"fmt"
	"sync"
)

// Global variables of Mutex and WaitGroup
var wg sync.WaitGroup
var mu sync.Mutex

// Initializing map globaly that works as 'Cashe'
var cashe map[string]bool = make(map[string]bool)

type Fetcher interface {

	// Fetch returns the body of URL and
	// a slice of URLs found on that page.

	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.

func Crawl(url string, depth int, fetcher Fetcher , count *int) {
	
	// If condition to prevent negative wait counter on the call from main 
	if(*count != 1){
		defer wg.Done()
	}
	
	// Locks  Cashe so other rotutines don't have conflict
	mu.Lock()
	tempFlag := cashe[url]
	mu.Unlock()

	// If url is already visited return
	if depth <= 0 || tempFlag {
		return
	}

	// Locks  Cashe so other rotutines don't have conflict
	mu.Lock()
	cashe[url] = true
	mu.Unlock()
	
	// Fetching Url 
	body, urls, err := fetcher.Fetch(url)

	// Prints error if there is any
	if err != nil {
		fmt.Println(err)
		return
	}

	// Url found
	fmt.Printf("found: %s %q\n", url, body)

	// Loops around all urls
	for _, u := range urls {
		
		// Adds wait counter
		wg.Add(1)
		*count++

		// Recursive call for every url
		go Crawl(u, depth-1, fetcher, count)
	}
}

func main() {

    count := 1
	Crawl("https://golang.org/", 4, fetcher, &count)
	
	// Waits until all go routines are not finshed
	wg.Wait()
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
