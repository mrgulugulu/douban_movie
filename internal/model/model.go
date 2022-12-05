package model

type DoubanMovie struct {
	Title      string `gorm:"title" form:"title" json:"title"`
	Subtitle   string `gorm:"subtitle" form:"subtitle" json:"subtitle"`
	Other      string `gorm:"other" form:"other" json:"other"`
	Url        string `gorm:"url" form:"url" json:"url"`
	Desc       string `gorm:"desc" form:"desc" json:"desc"`
	Year       string `gorm:"year" form:"year" json:"year"`
	Area       string `gorm:"area" form:"area" json:"area"`
	Tag        string `gorm:"tag" form:"tag" json:"tag"`
	Star       string `gorm:"star" form:"star" json:"star"`
	Comment    string `gorm:"comment" form:"comment" json:"comment"`
	ViewNumber int    `gorm:"view_number" form:"view_number" json:"view_number"`
	Quote      string `gorm:"quote" form:"quote" json:"quote"`
}

type Page struct {
	Page int
	Url  string
}
