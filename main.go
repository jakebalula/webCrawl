package main

import (
  "fmt"
  "golang.org/x/net/html"
  "io"
  "net/http"
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

func extractLinks(htmlContent string) []string {}

func main() {
  url := "https://github.com/jakebalula"
  content, err := fetchUrl(url)
  if err != nil {
    fmt.Println("Error fetching URL:", err)
    return
  }

  fmt.Println("Page content:", content[:500]) //This prints the first 500 characters
}
