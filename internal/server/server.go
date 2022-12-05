package server

import (
	"fmt"
	"net/http"

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

	fmt.Println("listening ", ser.Addr)
	ser.ListenAndServe()
}
