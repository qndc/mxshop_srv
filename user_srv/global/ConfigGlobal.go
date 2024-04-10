package global

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"mxshop_srvs/user_srv/config"
)

var Config = &config.Config{}

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("user-srv/config/config-dev.yml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("load config failed,reason:%s", err.Error()))
	}
	err = v.Unmarshal(Config)
	if err != nil {
		panic(fmt.Sprintf("unmarshal config failed,reason:%s", err.Error()))
	}

	//从Nacos加载业务配置
	IninNacosConfig()

	//动态监控配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		if e.Op&fsnotify.Write != 0 {
			Sugar.Infof("config file changed")
			_ = v.ReadInConfig()
			_ = v.Unmarshal(Config)
			//从Nacos加载业务配置
			IninNacosConfig()
		}
	})
}
