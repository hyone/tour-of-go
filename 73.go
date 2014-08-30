package main

import (
  "fmt"
)

type Fetcher interface {
  // Fetch returns the body of URL and
  // a slice of URLs found on that page.
  Fetch(url string) (body string, urls []string, err error)
}

type Result struct {
  url string
  body string
  urls []string
  err error
  depth int
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
  var workers = 0
  var visitedUrls = make(map[string]bool)

  ch   := make(chan *Result, 10)
  quit := make(chan bool, 5)

  workers++
  visitedUrls[url] = true
  go _Crawl(url, depth, fetcher, ch)

  for {
    select {
    case result := <-ch:
      showResult(result)
      for _, u := range result.urls {
        if d := result.depth-1; !visitedUrls[u] && d > 0  {
          visitedUrls[u] = true
          workers++
          go _Crawl(u, d, fetcher, ch)
        }
      }
      quit <- true
    case <-quit:
      if workers--; workers <= 0 {
        return
      }
    }
  }
}

func _Crawl(url string, depth int, fetcher Fetcher, ch chan *Result) {
  body, urls, err := fetcher.Fetch(url)
  ch <- &Result {
    url: url,
    body: body,
    urls: urls,
    err: err,
    depth: depth,
  }
}

func showResult(result *Result) {
  if result.err != nil {
    fmt.Println(result.err)
    return
  }

  fmt.Printf("found: %s %q\n", result.url, result.body)
}


func main () {
  Crawl("http://golang.org/", 4, fetcher)
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
var fetcher = fakeFetcher {
  "http://golang.org/": &fakeResult {
    "The Go Programming Language",
    []string {
      "http://golang.org/pkg/",
      "http://golang.org/cmd/",
    },
  },
  "http://golang.org/pkg/": &fakeResult{
      "Packages",
      []string{
          "http://golang.org/",
          "http://golang.org/cmd/",
          "http://golang.org/pkg/fmt/",
          "http://golang.org/pkg/os/",
      },
  },
  "http://golang.org/pkg/fmt/": &fakeResult{
      "Package fmt",
      []string{
          "http://golang.org/",
          "http://golang.org/pkg/",
      },
  },
  "http://golang.org/pkg/os/": &fakeResult{
      "Package os",
      []string{
          "http://golang.org/",
          "http://golang.org/pkg/",
      },
  },
}
