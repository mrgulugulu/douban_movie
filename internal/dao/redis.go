package dao

import (
	"encoding/json"
	"film-info/internal/model"
	"fmt"
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
func (d *Dao) WriteInRedis(filmInfo model.DoubanMovie) error {
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
