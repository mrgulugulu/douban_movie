package main

import (
	"film-info/config"
	"film-info/internal/dao"
	"log"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Printf("config error: %v", err)
	}
	movieTitleList, err := dao.D.GetMovieSetMembers(config.QueryMovieSet)
	if err != nil {
		log.Printf("get movie set member error: %v", err)
	}
	err = dao.D.CalViewNumber("千与千寻")
	for _, title := range movieTitleList {
		viewNum, err := dao.D.GetMovieViewNumber(title)
		if err != nil {
			log.Printf("get movie view number error: %v", err)
		}
		err = dao.D.UpdateMovieInfo(title, "ViewNumber", viewNum)
		if err != nil {
			log.Printf("update movie viewNumber error: %v", err)
		}
	}

}
