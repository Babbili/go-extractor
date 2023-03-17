package main

import (
	"fmt"
	"goget/getkeyword"
	"goget/geturl"
)

func main() {
	keyword := "Dubai"
	url := "https://landscape.cncf.io/members"

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
