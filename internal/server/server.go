package server

import (
	"context"
	"film-info/config"
	"film-info/internal/dao"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Addr string
	Port string
}

func (s *Server) Run() {
	r := gin.Default()

	r.GET("/filmInfo", query)
	r.GET("/filmInfo/top10", top10)
	r.DELETE("/filmInfo", delete)
	r.POST("/filmInfo", update)

	ser := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.Addr, s.Port),
		Handler: r,
	}
	// 搞个signal来监听，实现优雅关闭
	fmt.Println("listening ", ser.Addr)
	go ser.ListenAndServe()
	gracefulExitServer(ser)
}

// gracefulExitServer 实现优雅关闭，保存redis中必要的数据，如点击数
func gracefulExitServer(ser *http.Server) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	<-ch
	nowTime := time.Now()
	movieTitleList, err := dao.D.GetMovieSetMembers(config.QueryMovieSet)
	if err != nil {
		log.Printf("get movie set member error: %v", err)
	}
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = ser.Shutdown(ctx)
	if err != nil {
		log.Printf("shutdown error: %v", err)
	}
	fmt.Println("-----exited-----", time.Since(nowTime))
}
