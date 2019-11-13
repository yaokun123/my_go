package main

import (
	"github.com/gocolly/colly"
	"fmt"
)

func main()  {
	c := colly.NewCollector(
		//colly.AllowedDomains("news.baidu.com"),
		colly.UserAgent("Opera/9.80 (Windows NT 6.1; U; zh-cn) Presto/2.9.168 Version/11.50"),
	)

	/*为collector添加回调函数*/


	//请求发出之前调用
	c.OnRequest(func(request *colly.Request) {
		//Request头部设定
		request.Headers.Set("Host","baidu.com")
		request.Headers.Set("Connection","keep-alive")
		request.Headers.Set("Accept","*/*")
		request.Headers.Set("Origin","")
		request.Headers.Set("Referer","http://www.baidu.com")
		request.Headers.Set("Accept-Encoding","gzip, deflate")
		request.Headers.Set("Accept-Language","zh-CN, zh;q=0.9")

		fmt.Println("Visiting",request.URL)
	})

	//请求过程中出现Error时调用
	c.OnError(func(response *colly.Response, e error) {
		fmt.Println("something went wrong:",e)
	})

	//收到response后调用
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Visited",response.StatusCode)
	})

	// 如果收到的内容是HTML，就在onResponse执行后调用
	// 对响应的HTML元素处理
	c.OnHTML("title", func(element *colly.HTMLElement) {
		fmt.Println("title:",element.Text)
	})
	c.OnHTML("body", func(element *colly.HTMLElement) {
		// <div class="hotnews" alog-group="focustop-hotnews"> 下所有的a解析
		element.ForEach(".hotnews a", func(i int, e *colly.HTMLElement) {
			band := e.Attr("href")
			title := e.Text
			fmt.Printf("新闻 %d : %s - %s\n",i,title,band)
			//element.Request.Visit(band)
		})
	})
	/*c.OnHTML("a[href]", func(element *colly.HTMLElement) {
		//element.Request.Visit(element.Attr("href"))
		fmt.Println(element.Attr("href"))
	})
	c.OnHTML("tr td:nth-of-type(1)", func(element *colly.HTMLElement) {
		fmt.Println("First column of a table row:",element.Text)
	})*/

	//如果收到的内容是HTML或者XML，就在onHTML执行后调用
	c.OnXML("//h1", func(element *colly.XMLElement) {
		fmt.Println(element.Text)
	})

	//Scraped OnXML执行后调用
	c.OnScraped(func(response *colly.Response) {
		fmt.Println("Finished",response.Request.URL)
	})


	//入口
	c.Visit("http://news.baidu.com")
}
