package config

import "time"

// Redis信息
type redisConfig struct {
	Enable bool `yaml:"enable"`
	Addrs []string `yaml:"addrs"`
	Password string `yaml:"password"`
	DefaultDB int `yaml:"defaultDB"`
	DialTimeout time.Duration `yaml:"dialTimeout"`
}
