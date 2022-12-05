package main

import (
	"encoding/json"
	"film-info/config"
	"film-info/internal/dao"
	"film-info/internal/model"
	"fmt"
	"log"
	"time"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Printf("config error: %v", err)
	}
	film := model.DoubanMovie{Title: "aaa", Star: "10"}
	fileJson, _ := json.Marshal(film)
	cmd := dao.D.RedisDb.Set("filmaa", fileJson, time.Hour)
	filmaa, err := dao.D.RedisDb.Get("filma").Result()
	var film1 model.DoubanMovie
	json.Unmarshal([]byte(filmaa), &film1)
	fmt.Println(film1, cmd, err)

}
