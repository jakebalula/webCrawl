package main

import (
  "fmt"
  "golang.org/x/net/html"
  "io"
  "net/http"
  "strings"
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

func main() {
  url := "https://github.com/jakebalula"
  content, err := fetchUrl(url)
  if err != nil {
    fmt.Println("Error fetching URL:", err)
    return
  }
  links := extractLinks(content)
  fmt.Println("Extracted links:")
  for _, link := range links {
    fmt.Println(link)
  }
  fmt.Println("Page content:", content[:500]) //This prints the first 500 characters
}
