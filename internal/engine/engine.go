// engine 爬虫引擎
package engine

import (
	"bytes"
	"film-info/config"
	"film-info/internal/fetcher"
	"film-info/internal/model"
	"film-info/internal/parser"
	"strings"

	"log"

	"github.com/PuerkitoBio/goquery"
)

func Run(baseUrl string) []model.DoubanMovie {
	// 这里传进来的是base，接着就要进行拼接
	var movies []model.DoubanMovie
	pages := parser.GetPages(baseUrl)
	for _, page := range pages {
		url := strings.Join([]string{baseUrl, page.Url}, "")
		contents, err := fetcher.Fetch(url, config.Header)
		if err != nil {
			log.Print(err)
		}
		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(contents))
		if err != nil {
			log.Print(err)
		}
		movies = append(movies, parser.ParseMovies(doc)...)
	}
	return movies
}
