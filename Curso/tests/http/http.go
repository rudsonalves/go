package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
)

func main() {
	url := "https://youtube.com"
	resp, _ := http.Get(url)
	html, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(reflect.TypeOf(html))
	r, _ := regexp.Compile("<title>(.*?)</title>")
	fmt.Println(r)
	find := r.FindStringSubmatch(string(html))
	// fmt.Println(find)
	fmt.Println(find[1])
	// fmt.Println(string(html))
}
