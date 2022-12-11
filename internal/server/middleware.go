package server

import (
	"errors"
	"film-info/config"
	"film-info/internal/dao"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// recordUserVisit 限制同一ip访问次数
func recordUserVisit(c *gin.Context) {
	curIp := c.RemoteIP()
	curUVStr, err := dao.D.Get(curIp)
	curUV, _ := strconv.Atoi(curUVStr)
	if !errors.Is(err, redis.Nil) && curUV > config.UserVisitLimit {
		c.AbortWithStatusJSON(http.StatusBadRequest, "sorry, your visit is too frequent!")
	} else {
		val, err := dao.D.Incr(curIp)
		if err != nil {
			log.Printf("incr ip error: %v", err)
		}
		if val == 1 {
			err = dao.D.Expire(curIp, config.Minute)
			if err != nil {
				log.Printf("expire ip error: %v", err)
			}
		}
	}
}
