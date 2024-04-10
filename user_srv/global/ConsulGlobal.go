package global

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
)

// InitAndRegister 服务注册+开启健康检查
func InitAndRegister() {
	//初始化客户端，配置服务端地址，consul服务为服务端
	cfg := api.DefaultConfig()
	consulAddr := fmt.Sprintf("%s:%d", Config.ConsulInfo.Host, Config.ConsulInfo.Port)
	cfg.Address = consulAddr
	client, _ := api.NewClient(cfg)

	//生成检查对象，用于健康检查，会开启一个协程定时访问指定Url
	check := new(api.AgentServiceCheck)
	//check.CheckID = id                                           // 健康检查项的id，唯一
	//check.Name = name                                            // 检查项的名字
	chkUrl := fmt.Sprintf("%s:%d", Config.ConsulInfo.Host, Config.ServerPort)
	Sugar.Infof("服务检查地址：%s", chkUrl)
	check.GRPC = chkUrl   // 定期访问的Url,通过这个url请求结果确定服务是否正常
	check.Timeout = "5s"  // 请求超时时间，5秒
	check.Interval = "5s" // 定期检查的时间间隔，这里是5秒
	check.Method = "GET"  // 设置http请求方式，默认是GET
	check.DeregisterCriticalServiceAfter = "10s"
	//check.Header // 可以自定义请求头，可以不配置

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.ID = fmt.Sprintf("%s", uuid.NewV4())
	registration.Port = Config.ServerPort
	registration.Address = Config.ConsulInfo.Host
	registration.Name = Config.ServerName
	registration.Tags = []string{"frp", "grpc"}
	registration.Check = check

	client.Agent().ServiceRegister(registration)

	//优雅退出，退出是注销已注册的服务，监听退出事件:ctrl+c
	//quit := make(chan os.Signal)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//ServiceDeregister(registration.ID)
}

// GetAllServices 服务发现，获取所有的服务
func GetAllServices() {
	//初始化客户端，配置服务端地址，consul服务为服务端
	cfg := api.DefaultConfig()
	cfg.Address = "119.91.238.152:8500"
	client, _ := api.NewClient(cfg)
	services, _ := client.Agent().Services()
	for key, _ := range services {
		fmt.Println(key)
	}
}

// FilterService 获取服务并对其进行过滤
func FilterService() {
	//初始化客户端，配置服务端地址，consul服务为服务端
	cfg := api.DefaultConfig()
	cfg.Address = "119.91.238.152:8500"
	client, _ := api.NewClient(cfg)
	service, _ := client.Agent().ServicesWithFilter(`Service == "user-web"`)
	fmt.Println(service)

}

// ServiceDeregister 注销已注册的服务
func ServiceDeregister(serviceId string) {
	//初始化客户端，配置服务端地址，consul服务为服务端
	cfg := api.DefaultConfig()
	consulAddr := fmt.Sprintf("%s:%d", Config.ConsulInfo.Host, Config.ConsulInfo.Port)
	cfg.Address = consulAddr
	client, _ := api.NewClient(cfg)
	err := client.Agent().ServiceDeregister(serviceId)
	if err != nil {
		Sugar.Errorf("服务注销失败:%s", err.Error())
	} else {
		Sugar.Info("服务注销成功")
	}
}
