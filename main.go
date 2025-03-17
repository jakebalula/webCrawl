package main

import (
  "fmt"
  "golang.org/x/net/html"
  "io"
  "net/http"
  "strings"
  "sync"
)

// fetchURL sends an HTTP Get request and returns the response body as a string.
func fetchUrl(url string) (string, error) {
  resp, err := http.Get(url)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close() //Close the response body

  body, err := io.ReadAll(resp.Body) //Read the response body
  if err != nil {
    return "", err
  }
  return string(body), nil
}

func extractLinks(htmlContent string) []string {
  var links []string
  tokenizer := html.NewTokenizer(strings.NewReader(htmlContent))
  for {
    tt := tokenizer.Next()
    switch tt {
    case html.ErrorToken:
      return links
    case html.StartTagToken:
      token := tokenizer.Token()
      if token.Data == "a" {
        for _, attr := range token.Attr {
          if attr.Key == "href" {
            links = append(links, attr.Val)
          }
        }
      }
    }
  }
}

func crawl(url string, depth int, visited map[string]bool, mu *sync.Mutex, wg *sync.WaitGroup) {
  defer wg.Done()

  if depth <= 0 {
    return
  }
  mu.Lock()
  if visited[url] {
    mu.Unlock()
    return
  }
  visited[url] = true
  mu.Unlock()

  fmt.Println("Crawling:", url)

  content, err := fetchUrl(url)
  if err != nil {
    fmt.Println("Error fetching:", url)
    return
  }
  links := extractLinks(content)
  for _, link := range links {
    wg.Add(1)
    go crawl(link, depth-1, visited, mu, wg)
  }
}

func main() {
  url := "https://github.com/jakebalula"
  maxDepth := 5
  var wg sync.WaitGroup
  var mu sync.Mutex
  visited := make(map[string]bool)

  wg.Add(1)
  go crawl(url, maxDepth, visited, &mu, &wg)
  wg.Wait()
  fmt.Println("Crawled!")
}
