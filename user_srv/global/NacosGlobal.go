package global

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func IninNacosConfig() {
	// 创建serverConfig
	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      Config.NacosInfo.IpAddr,
			ContextPath: "/nacos",
			Port:        Config.NacosInfo.Port,
			Scheme:      "http",
		},
	}
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         Config.NacosInfo.NamespaceId, // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              Config.NacosInfo.LogDir,
		CacheDir:            Config.NacosInfo.CacheDir,
		LogLevel:            Config.NacosInfo.LogLevel,
	}
	// 创建动态配置客户端
	ConfigClient, ccerr := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if ccerr != nil {
		panic("与配置中心建立连接失败")
	}
	configContant, gcerr := ConfigClient.GetConfig(vo.ConfigParam{
		DataId: Config.NacosInfo.DataId,
		Group:  Config.NacosInfo.Group})
	if gcerr != nil {
		panic("从配置中心读取用户服务配置失败")
	}

	//将读取的配置文件解析为struct
	//注意：想要将json字符串内容解析为struct对象，需要在struct中配置Tag,如下所示：
	// Host string `json:"host"`
	userr := json.Unmarshal([]byte(configContant), Config)
	if userr != nil {
		panic("解析用户服务配置失败")
	}

	//监听Nacos配置中心用户服务配置文件变化
	lcerr := ConfigClient.ListenConfig(vo.ConfigParam{
		DataId: Config.NacosInfo.DataId,
		Group:  Config.NacosInfo.Group,
		OnChange: func(namespace, group, dataId, data string) {
			//Sugar.Infof("监听到配置文件变化：group:%s , dataId:%s , data:\r\n%s", group, dataId, data)
			//将变更后的配置文件解析为struct
			json.Unmarshal([]byte(configContant), Config)
		},
	})
	if lcerr != nil {
		panic("监听用户配置异常")
	}
}
