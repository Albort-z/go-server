package conf

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Addr string `json:"addr"`
}

func ReadConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigFile("config/config.json")

	v.OnConfigChange(func(in fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", in.Name)
		log.Println(v.Get("port"))
		// 在这里触发服务平滑更新
	})
	v.WatchConfig()

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var c Config
	err = v.Unmarshal(&c)
	return &c, err
}
