package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Cocurrency Patterns

// <-chan - canal somente leitura
func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)</title>")
			result := r.FindStringSubmatch(string(html))
			if len(result) > 1 {
				c <- result[1]
			} else {
				c <- "URL error: " + url
			}
		}(url) // necessário para passar a url para a função implícita
	}
	return c
}

func main() {
	urls := []string{
		"https://diolinux.com.br",
		"https://dbikeshop.com.br",
		"https://www.cod3r.com.br",
		"https://google.com",
		"https://github.com",
		"https://youtube.com",
		"https://www.udemy.com/"}

	c := titulo(urls...)
	fmt.Println("Em ordem de resposta:")
	for i, _ := range urls {
		fmt.Printf("%d: ", i+1)
		fmt.Print(<-c, "\n")
	}
}
