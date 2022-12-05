package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

var (
	BaseUrl = "https://movie.douban.com/top250"
	Header  = map[string]string{
		"Host":                      "movie.douban.com",
		"Connection":                "keep-alive",
		"Cache-Control":             "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Referer":                   "https://movie.douban.com/top250",
	}
)

const (
	QueryMovieSet    = "movieset"
	QueryExpiredTime = time.Hour
	ViewNumber       = "viewnum"
)

type MysqlConfig struct {
	MysqlIP   string `mapstructure:"ip"`
	MysqlPort string `mapstructure:"port"`
	MysqlUser string `mapstructure:"user"`
	MysqlPwd  string `mapstructure:"pwd"`
	DataBase  string `mapstructure:"database"`
}

type RedisConfig struct {
	RedisIP   string `mapstructure:"ip"`
	RedisPort string `mapstructure:"port"`
	RedisPwd  string `mapstructure:"pwd"`
	DataBase  int    `mapstructure:"db"`
}
type Config struct {
	Mysqlcfg  *MysqlConfig  `mapstructure:"mysql"`
	Rediscfg  *RedisConfig  `mapstructure:"redis"`
	ServerCfg *ServerConfig `mapstructure:"server"`
}

type ServerConfig struct {
	Addr string `mapstructure:"addr"`
	Port string `mapstructure:"port"`
}

var ServiceConf *Config

func LoadConfig(confFile ...string) error {
	c := viper.New()
	conf := Config{}
	c.AddConfigPath("../config")
	c.AddConfigPath("../config")
	c.SetConfigName("config")
	c.SetConfigType("yaml")
	err := c.ReadInConfig()
	if err != nil {
		return fmt.Errorf("read config error: %v", err)
	}
	err = c.Unmarshal(&conf)
	ServiceConf = &conf
	return err
}
