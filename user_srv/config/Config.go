package config

type Config struct {
	ServerName string       `json:"server-name"`
	ServerHost string       `json:"server-host"`
	ServerPort int          `json:"server-port"`
	MySqlInfo  MySqlConfig  `json:"mysql"`
	ConsulInfo ConsulConfig `json:"consul"`
	NacosInfo  NacosConfig  `mapstructure:"nacos"`
}

type MySqlConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ConsulConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type NacosConfig struct {
	IpAddr      string `mapstructure:"host"`
	Port        uint64 `mapstructure:"port"`
	NamespaceId string `mapstructure:"namespace"`
	LogDir      string `mapstructure:"logDir"`
	CacheDir    string `mapstructure:"cacheDir"`
	LogLevel    string `mapstructure:"logLevel"`
	DataId      string `mapstructure:"dataId"`
	Group       string `mapstructure:"group"`
}
