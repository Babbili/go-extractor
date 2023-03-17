## Golang keyword getter
> https://pkg.go.dev

simple implementation of a function in Golang that takes a URL as input and checks if the webpage contains the keyword "cat"

```go
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
)

func containsCat(url string) bool {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error: ", err)
        return false
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error: ", err)
        return false
    }

    return strings.Contains(string(body), "cat")
}

func main() {
    url := "https://www.example.com"
    result := containsCat(url)
    fmt.Println(result)
}
```

function in Golang that uses the `net/http` and `golang.org/x/net/html` packages to extract all external links from a given webpage

```go
package main

import (
    "fmt"
    "net/http"
    "strings"

    "golang.org/x/net/html"
)

func getExternalLinks(url string) ([]string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    links := []string{}
    tokenizer := html.NewTokenizer(resp.Body)
    for {
        tokenType := tokenizer.Next()
        if tokenType == html.ErrorToken {
            return links, nil
        }
        if tokenType == html.StartTagToken {
            token := tokenizer.Token()
            if token.Data == "a" {
                for _, attr := range token.Attr {
                    if attr.Key == "href" {
                        link := attr.Val
                        if strings.HasPrefix(link, "http") && !strings.HasPrefix(link, url) {
                            links = append(links, link)
                        }
                    }
                }
            }
        }
    }
}

func main() {
    url := "https://example.com"
    externalLinks, err := getExternalLinks(url)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("External Links in %v:\n", url)
        for _, link := range externalLinks {
            fmt.Println(link)
        }
    }
}
```

This function takes a URL as an argument, sends an HTTP GET request to that URL, and parses the HTML response to extract all external links (<a> tags with an href attribute starting with "http" but not starting with the given URL). It returns a slice of strings containing the URLs of all external links found in the webpage. If there is an error in sending the HTTP GET request or parsing the HTML, the function returns an error

<br />

```go
package main

import (
    "fmt"
    "net/http"
    "golang.org/x/net/html"
    "strings"
    "regexp"
)

// Extracts all external links from the given webpage URL
func extractExternalLinks(url string) ([]string, error) {
    // Send HTTP GET request to the URL
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Parse HTML document
    doc, err := html.Parse(resp.Body)
    if err != nil {
        return nil, err
    }

    // Extract links recursively
    links := make([]string, 0)
    var extract func(*html.Node)
    extract = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            for _, attr := range n.Attr {
                if attr.Key == "href" {
                    href := attr.Val
                    // Check if link is external
                    if isExternalLink(href) {
                        links = append(links, href)
                    } else if isInternalLink(href) {
                        // Extract links from internal page
                        internalLinks, err := extractExternalLinks(href)
                        if err == nil {
                            links = append(links, internalLinks...)
                        }
                    }
                    break
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            extract(c)
        }
    }
    extract(doc)

    return links, nil
}

// Checks if the given link URL is external
func isExternalLink(url string) bool {
    return strings.HasPrefix(url, "http") || strings.HasPrefix(url, "https")
}

// Checks if the given link URL is internal to the same domain
func isInternalLink(url string) bool {
    re := regexp.MustCompile(`^\/[^\/].*$`)
    return re.MatchString(url)
}

func main() {
    links, err := extractExternalLinks("https://www.example.com")
    if err != nil {
        fmt.Println(err)
    } else {
        for _, link := range links {
            fmt.Println(link)
        }
    }
}
```
This function sends an HTTP GET request to the given URL, parses the HTML document using the html package, and recursively extracts all <a> tags with an href attribute. If the link is external, it is added to the links slice. If the link is internal, the function is called recursively to extract external links from the internal page, and those links are added to the links slice.

The isExternalLink function checks if the given link URL is external by checking if it starts with "http" or "https". The isInternalLink function checks if the given link URL is internal by using a regular expression to match URLs that start with a forward slash ("/") followed by any character except another forward slash. This assumes that all internal links on the same domain start with a forward slash and do not contain a protocol (e.g. "http", "https", "ftp", etc.)

