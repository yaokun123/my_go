package main

import (
	"net/url"
	"log"
	"github.com/gocolly/colly"
	"time"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"bytes"
)

func main()  {
	urlstr := "https://news.baidu.com"
	u,err := url.Parse(urlstr)
	if err != nil {
		log.Fatal(err)
	}

	c := colly.NewCollector()
	//超时设定
	c.SetRequestTimeout(100*time.Second)
	//指定agent信息
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.108 Safari/537.36"

	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("Host", u.Host)
		request.Headers.Set("Connection", "keep-alive")
		request.Headers.Set("Accept", "*/*")
		request.Headers.Set("Origin", u.Host)
		request.Headers.Set("Referer", urlstr)
		request.Headers.Set("Accept-Encoding", "gzip, deflate")
		request.Headers.Set("Accept-Language", "zh-CN, zh;q=0.9")
	})

	c.OnHTML("title", func(element *colly.HTMLElement) {
		fmt.Println("title",element.Text)
	})

	c.OnResponse(func(response *colly.Response) {
		//goquery直接读取response.Body的内容
		htmlDoc,_ := goquery.NewDocumentFromReader(bytes.NewReader(response.Body))

		//读取url再传给goquery，访问url读取内容，此处不建议使用
		//htmlDoc,_ := goquery.NewDocument(response.Request.URL.String())

		// 找到抓取项 <div class="hotnews" alog-group="focustop-hotnews"> 下所有的a解析
		htmlDoc.Find(".hotnews a").Each(func(i int, selection *goquery.Selection) {
			band, _ := selection.Attr("href")
			title := selection.Text()
			fmt.Printf("热点新闻 %d: %s - %s\n", i, title, band)
		})

	})

	c.Visit(urlstr)

}
