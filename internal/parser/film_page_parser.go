package parser

import (
	"film-info/internal/model"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// 分析电影数据
func ParseMovies(doc *goquery.Document) (movies []model.DoubanMovie) {
	doc.Find("ol.grid_view > li > div.item").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".hd a span").Eq(0).Text()
		url, _ := s.Find("a").Attr("href")
		subtitle := s.Find(".hd a span").Eq(1).Text()
		subtitle = strings.TrimLeft(subtitle, "  / ")

		other := s.Find(".hd a span").Eq(2).Text()
		other = strings.TrimLeft(other, "  / ")

		desc := strings.TrimSpace(s.Find(".bd p").Eq(0).Text())
		DescInfo := strings.Split(desc, "\n")
		desc = DescInfo[0]

		movieDesc := strings.Split(DescInfo[1], "/")
		year := strings.TrimSpace(movieDesc[0])
		area := strings.TrimSpace(movieDesc[1])
		tag := strings.TrimSpace(movieDesc[2])

		star := s.Find(".bd .star .rating_num").Text()

		comment := strings.TrimSpace(s.Find(".bd .star span").Eq(3).Text())
		compile := regexp.MustCompile("[0-9]")
		comment = strings.Join(compile.FindAllString(comment, -1), "")

		quote := s.Find(".quote .inq").Text()

		movie := model.DoubanMovie{
			Title:    title,
			Subtitle: subtitle,
			Other:    other,
			Url:      url,
			Desc:     desc,
			Year:     year,
			Area:     area,
			Tag:      tag,
			Star:     star,
			Comment:  comment,
			Quote:    quote,
		}
		movies = append(movies, movie)
	})

	return movies
}
