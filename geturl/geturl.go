package geturl

import (
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// Extracts all external links from the given webpage URL
func ExtractExternalLinks(url string) ([]string, error) {
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
						internalLinks, err := ExtractExternalLinks(href)
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
