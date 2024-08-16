package config

import "time"

// Redis信息
type redisConfig struct {
	Enable bool `yaml:"enable"`
	Addrs []string `yaml:"addrs"`
	Password string `yaml:"password"`
	Username string `yaml:"username"`
	DefaultDB int `yaml:"defaultDB"`
	DialTimeout time.Duration `yaml:"dialTimeout"`
}
