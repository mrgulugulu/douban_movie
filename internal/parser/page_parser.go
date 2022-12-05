package parser

import (
	"bytes"
	"film-info/config"
	"film-info/internal/fetcher"
	"film-info/internal/model"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// ParsePages 解析主页的内容，并将解析到的所有电影url保存下来
func ParsePages(doc *goquery.Document) []model.Page {
	var pages []model.Page
	pages = append(pages, model.Page{Page: 1, Url: ""})
	doc.Find(".paginator > a").Each(func(i int, s *goquery.Selection) {
		page, _ := strconv.Atoi(s.Text())
		url, _ := s.Attr("href")

		pages = append(pages, model.Page{
			Page: page,
			Url:  url,
		})
	})

	return pages
}

func GetPages(url string) []model.Page {
	contents, err := fetcher.Fetch(url, config.Header)
	if err != nil {
		log.Panicln(err)
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(contents))
	if err != nil {
		log.Fatal(err)
	}
	res := ParsePages(doc)

	return res
}
