package dao

import (
	"encoding/json"
	"film-info/config"
	"film-info/internal/model"
	"fmt"
	"log"
	"time"
)

// 读写的逻辑分别是
// 读：先读redis，没再读mysql。这里返回的应该就是list
func (d *Dao) ReadFromRedis(title string) (model.DoubanMovie, error) {
	var result model.DoubanMovie
	filmInfo, err := d.RedisDb.Get(title).Result()
	if err != nil {
		return result, fmt.Errorf("data not found in redis: %v", err)
	}
	err = json.Unmarshal([]byte(filmInfo), &result)
	if err != nil {
		return result, fmt.Errorf("unmarshal error: %v", err)
	}
	return result, nil
}

// 写：先更新mysql，再删redis
func (d *Dao) WriteFilmInfoInRedis(filmInfo model.DoubanMovie) error {
	filmJson, err := json.Marshal(filmInfo)
	if err != nil {
		return fmt.Errorf("filmInfo marshal error: %v", err)
	}
	cmdRes := d.RedisDb.Set(filmInfo.Title, filmJson, time.Minute)
	if cmdRes.Err() != nil {
		return fmt.Errorf("write in redis error: %v", cmdRes.Err())
	}
	return nil
}

// redis删除操作
func (d *Dao) DelFromRedis(filmInfo model.DoubanMovie) error {
	filmJson, err := json.Marshal(filmInfo)
	if err != nil {
		return fmt.Errorf("filmInfo marshal error: %v", err)
	}
	cmdRes := d.RedisDb.Del(filmInfo.Title, string(filmJson))
	if cmdRes.Err() != nil {
		return fmt.Errorf("delete data error: %v", err)
	}
	return nil
}

// CalViewNumber 统计记录的浏览数
func (d *Dao) CalViewNumber(filter string) error {
	// 先检查是否已经有了，没就set，有就+1
	res, err := d.RedisDb.HIncrBy("\""+filter+"\"", config.ViewNumber, 1).Result()
	if err != nil {
		return fmt.Errorf("incr error: %v", err)
	}
	log.Print(res)
	return nil
}

// QueryMovieSetAdd 添加movie的title进set中
func (d *Dao) QueryMovieSetAdd(filter string) {
	_, err := d.RedisDb.SAdd(config.QueryMovieSet, filter).Result()
	if err != nil {
		log.Printf("set add error: %v", err)
	}
}

// QueryMovieSetMembers 返回set中的movie title
func (d *Dao) GetMovieSetMembers(filter string) ([]string, error) {
	res, err := d.RedisDb.SMembers(filter).Result()
	if err != nil {
		return nil, fmt.Errorf("redis queryMovieSet members error %v", err)
	}
	return res, nil
}

// GetMovieViewNumber 返回redis中电影的view number
func (d *Dao) GetMovieViewNumber(movieTtile string) (string, error) {
	res, err := d.RedisDb.HGet("\""+movieTtile+"\"", config.ViewNumber).Result()
	if err != nil {
		return "", fmt.Errorf("GetMovieViewNumber error %v", err)
	}
	return res, nil
}
