package gohtml

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

//Get the title from an html page
func Title(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile(`<title>(.*?)<\/title>`)
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}
