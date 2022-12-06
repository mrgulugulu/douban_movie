package server

import (
	"film-info/internal/dao"
	"film-info/internal/model"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 这里补上redis的逻辑
func query(c *gin.Context) {
	// filmTitle := c.DefaultQuery("title", "")
	// 这里使用了shouldbind来通过反射匹配，但好似有点损失性能
	var filmInfo model.DoubanMovie
	if err := c.ShouldBind(&filmInfo); err == nil {
		// 写入set中
		dao.D.QueryMovieSetAdd(filmInfo.Title)
		err := dao.D.CalViewNumber(filmInfo.Title)
		if err != nil {
			log.Printf("cal viewNumber error %v", err)
		}
		redisRes, err := dao.D.ReadFromRedis(filmInfo.Title)
		if err == nil {
			c.String(http.StatusOK, fmt.Sprintf("%+v", redisRes))
			return
		}
		info, err := dao.D.QueryMovieInfo(filmInfo.Title)
		if err != nil {
			c.String(http.StatusNotFound, fmt.Sprintf("%+v", err))
		} else {
			// 这里要写入redis
			_ = dao.D.WriteFilmInfoInRedis(info)
			c.String(http.StatusOK, fmt.Sprintf("%+v", info))
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

//
func top10(c *gin.Context) {
	result := make([]model.DoubanMovie, 0, 10)
	filter := c.DefaultQuery("column", "")
	// 可以先检查redis中是否存在该set
	if dao.D.IsZSetExist(filter) {
		movieList := dao.D.ReadZSet(filter)
		for _, movieTitle := range movieList {
			info, err := dao.D.ReadFromRedis(movieTitle)
			if err != nil {
				log.Printf("read from redis error: %v", err)
			}
			result = append(result, info)
		}
		c.String(http.StatusOK, fmt.Sprintf("%v+", result))
		return
	}
	res, err := dao.D.Top10(filter)
	dao.D.SetTop10ZSet(filter, res)
	for _, movieInfo := range res {
		err := dao.D.WriteFilmInfoInRedis(movieInfo)
		if err != nil {
			log.Printf("write info in redis error: %v", err)
		}
	}
	if err != nil {
		c.String(http.StatusNotFound, fmt.Sprintf("%+v", err))
	} else {
		c.String(http.StatusOK, fmt.Sprintf("%+v", res))
	}
}

func delete(c *gin.Context) {
	// 先查redis是否有，有再删除，没得话直接删mysql
	filmTitle := c.DefaultQuery("title", "")
	redisRes, err := dao.D.ReadFromRedis(filmTitle)
	if err == nil {
		_ = dao.D.DelFromRedis(redisRes)
	}
	err = dao.D.DeleteMovieInfo(filmTitle)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("delete data error: %v", err))
	} else {
		c.String(http.StatusOK, "data delete successfully")
	}
}

// 修改电影的信息
func update(c *gin.Context) {
	// 先修改mysql，再删redis
	var movieInfo model.DoubanMovie

	title := c.PostForm("title")
	key, value := c.PostForm("key"), c.PostForm("value")
	movieInfo, err := dao.D.QueryMovieInfo(title)
	if err != nil {
		c.String(http.StatusNotFound, fmt.Sprintf("update error: %v", err))
	}
	err = dao.D.UpdateMovieInfo(title, key, value)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("update error: %v", err))
		return
	} else {
		c.String(http.StatusOK, "update data successfully")
	}
	_ = dao.D.DelFromRedis(movieInfo)
}
