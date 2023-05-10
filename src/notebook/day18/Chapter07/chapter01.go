package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//定义结构体Movie，保存数据信息
type Movie struct {
	Name        string
	Star        string
	Releasetime string
	Score       string
}

//
func get_data(offest string) string {
	//urls := "https://www.maoyan.com/board/4?offset=" + offest
	urls := "https://www.douyu.com/g_yz"
	fmt.Println(urls)

	//定义请求对象NewRequest
	req, _ := http.NewRequest("GET", urls, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.56")
	transport := &http.Transport{}

	//在Client设置参数Transport即可实现代理IP
	client := &http.Client{Transport: transport}
	//发送http请求
	resp, _ := client.Do(req)
	//获取网站响应内容
	body, _ := ioutil.ReadAll(resp.Body)
	//网页响应内容转码
	result := string(body)
	//设置延时，请求过快会引发反爬
	time.Sleep(2 * time.Second)
	return result
}

func clean_data(data string) []map[string]string {
	//使用goquery解析HTML
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	fmt.Println("error:", err)
	//定义变量result和info
	var result []map[string]string
	var info map[string]string
	//遍历网页所有信息
	fmt.Println("begin")
	selection := dom.Find(".DyListCover-superscript")
	selection.Each(func(i int, selection *goquery.Selection) {
		//记录每一步电影，每存一部电影必须清空集合
		info = map[string]string{}
		name := selection.Find(".img").Text()
		//star := selection.Find(".DyListCover-intro").Text()
		//releasetime := selection.Find(".DyListCover-userName").Text()
		//score := selection.Find(".DyListCover-hotIcon").Text()
		//fmt.Println(name, star, releasetime, score)

		fmt.Println(name)
		info["name"] = strings.TrimSpace(name)
		//info["star"] = strings.TrimSpace(star)
		//info["releasetime"] = strings.TrimSpace(releasetime)
		//info["score"] = strings.TrimSpace(score)

		//将电影信息写入切片
		result = append(result, info)

	})
	fmt.Println("ending")
	return result
}

func main() {

	webData := get_data(strconv.Itoa(10))
	cleanData := clean_data(webData)
	if cleanData != nil {
		fmt.Println("Out:success")
	}

}
