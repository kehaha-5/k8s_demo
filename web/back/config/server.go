package config

// 应用信息
type app struct {
	Desc       string `yaml:"desc"`
	Addr       string `yaml:"addr"`
	ConfigFile string `yaml:"configFile"`
	Version    string `yaml:"version"`
	Env        string `yaml:"env"`
	TimeFormat string `yaml:"TimeFormat"`
}

// ServerConfig 配置信息
type ServerConfig struct {
	App   app         `yaml:"app"`
	Mysql mysqlConfig `yaml:"mysql"`
	Redis redisConfig `yaml:"redis"`
}

var Config ServerConfig
