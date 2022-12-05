package dao

import (
	"errors"
	"film-info/internal/model"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

// SaveMovieInfo 写入mysql数据库中
func (d *Dao) SaveMovieInfo(moviesInfo []model.DoubanMovie) int {
	count := 0
	for index, movie := range moviesInfo {
		// 写入之前先查询是否存在
		if _, err := d.QueryMovieInfo(movie.Title); err == nil {
			log.Printf("电影: %s 已存在", movie.Title)
			continue
		}
		if err := d.MysqlDb.Create(&movie).Error; err != nil {
			log.Printf("db.Create index: %d, err : %v", index, err)
			continue
		}
		count++
	}
	return count
}

// QueryMovieInfo 模糊查询电影信息
func (d *Dao) QueryMovieInfo(filter string) (model.DoubanMovie, error) {
	var result model.DoubanMovie
	filmName := filter

	if res := d.MysqlDb.Where("title LIKE ?", "%"+filmName+"%").First(&result); errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return result, errors.New("data not found")
	}
	return result, nil
}

// Top10 返回指定列的top10作品
func (d *Dao) Top10(filter string) ([]model.DoubanMovie, error) {
	var results []model.DoubanMovie
	if res := d.MysqlDb.Order(fmt.Sprintf("%s desc", filter)).Limit(10).Find(&results); errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("data not found")
	}
	return results, nil
}

// DeleteMovieInfo 使用模糊查询title删除电影记录
func (d *Dao) DeleteMovieInfo(filter string) error {
	// 先查询出记录的主键，再删除，精确删除
	var movie model.DoubanMovie
	filmInfo, err := d.QueryMovieInfo(filter)
	if err != nil {
		return fmt.Errorf("delete data error: %v", err)
	}
	res := d.MysqlDb.Where("title = ?", filmInfo.Title).Delete(&movie)
	if res.Error != nil {
		return fmt.Errorf("delete data error: %v", res.Error)
	}
	return nil
}

// UpdateMovieInfo 修改电影信息
func (d *Dao) UpdateMovieInfo(title, key, value string) error {
	var movieInfo model.DoubanMovie
	_, err := d.QueryMovieInfo(title)
	if err != nil {
		return fmt.Errorf("update info error: %v", err)
	}
	if res := d.MysqlDb.Model(&movieInfo).Where("title = ?", title).Update(key, value); res.Error != nil {
		return fmt.Errorf("mysql update error: %v", err)
	}
	return nil
}
