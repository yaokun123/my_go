package main

import (
	"net/url"
	"log"
	"github.com/gocolly/colly"
	"time"
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"fmt"
	"github.com/axgle/mahonia"
)

func main()  {
	urlstr := "http://www.dedecms.com/"
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

	/*c.OnHTML("title", func(element *colly.HTMLElement) {
		fmt.Println("title",element.Text)
	})*/

	c.OnResponse(func(response *colly.Response) {
		//goquery直接读取response.Body的内容
		htmlDoc,_ := goquery.NewDocumentFromReader(bytes.NewReader(response.Body))

		fmt.Println(GbkToUtf8(htmlDoc.Find("title").Text()))
	})

	c.Visit(urlstr)

}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func GbkToUtf8(src string)string{
	return ConvertToString(src, "gbk", "utf-8")
}