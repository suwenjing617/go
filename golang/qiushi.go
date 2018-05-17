package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func GetJokes() {
	// doc, err := goquery.NewDocument("https://www.jinse.com/xinwen")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// doc.Find(".js-main").Each(func(i int, s *goquery.Selection) {
	// 	// fmt.Println(s.Text())
	// 	con := s.Text()
	// 	fmt.Println(con)
	// })


	Name:        "金色财经",
	Description: "金色财经新闻 [https://www.jinse.com/xinwen]",
	// Pausetime: 300,
	// Keyin:     KEYIN,
	// Limit:     LIMIT,
	EnableCookie: false,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			ctx.AddQueue(&request.Request{Url: "https://www.jinse.com/xinwen", Rule: "获取新闻URL"})
		},
		Trunk: map[string]*Rule{

			"获取新闻URL": {
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					list := query.Find(".article-main div ol a")
					list.Each(func(i int, s *goquery.Selection) {
						if i == 0 {
							return
						}

						if url, ok := s.Attr("href"); ok {
							fmt.Print("url is: ", url, "\n")

							ctx.AddQueue(&request.Request{
								Url:  url,
								Rule: "新闻详情",
								// Header: http.Header{"Content-Type": []string{"application/xml"}},
								Temp: map[string]interface{}{
									"url": url,
								},
							})
						}
					})
				},
			},
			"新闻详情": {
				ItemFields: []string{ // 设定了 ItemFields, 才会写入文件
					"title",
					"content",
					"url",
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()

					title := query.Find("h2").Text()
					content := query.Find("p").Text()
					// info := query.Find(".article-info").Text()

					fmt.Print("title: ", title, "\n")

					ctx.Output(map[int]interface{}{
						0: title,
						1: content,
						2: ctx.GetTemp("url", ""),
					})
				},
			},
		},
	},
}

func main() {
	GetJokes()
}
