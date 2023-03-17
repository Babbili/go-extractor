package getkeyword

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ContainsKeyword(url string, keyword string) bool {
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

	return strings.Contains(string(body), keyword)
}
