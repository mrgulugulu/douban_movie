package main

import (
	"film-info/config"
	"film-info/internal/server"
)

func main() {
	s := &server.Server{
		Addr: config.ServiceConf.ServerCfg.Addr,
		Port: config.ServiceConf.ServerCfg.Port,
	}

	// movies := engine.Run(config.BaseUrl)
	// insertedNumber := dao.D.SaveMovieInfo(movies)
	// fmt.Printf("insert %d data successfully", insertedNumber)
	s.Run()

}
