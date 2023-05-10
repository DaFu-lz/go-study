package Crawl

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Demo1(url string, ch chan<- string) {
	get, _ := http.Get(url)
	//fmt.Println(get)
	//fmt.Println(err)

	body, _ := ioutil.ReadAll(get.Body)
	fmt.Println(body)
	ch <- string(body)
}
