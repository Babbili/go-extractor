package main

import (
	"fmt"
	"goget/getkeyword"
	"goget/geturl"
)

func main() {
	keyword := "\b[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}\b"
	url := "https://landscape.cncf.io/"

	links, err := geturl.ExtractExternalLinks(url)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, link := range links {
			// fmt.Println(link)
			result := getkeyword.ContainsKeyword(url, keyword)
			if result {
				fmt.Println(result, link)
			}
		}
	}
}
