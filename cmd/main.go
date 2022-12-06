package main

import (
	"film-info/config"
	"film-info/internal/dao"
	"film-info/internal/engine"
	"film-info/internal/server"
	"fmt"
)

func main() {
	s := &server.Server{
		Addr: config.ServiceConf.ServerCfg.Addr,
		Port: config.ServiceConf.ServerCfg.Port,
	}

	movies := engine.Run(config.BaseUrl)
	insertedNumber := dao.D.SaveMovieInfo(movies)
	fmt.Printf("insert %d data successfully", insertedNumber)
	s.Run()

}
