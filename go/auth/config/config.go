package config

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

// Config defines the configuration structure.
type Config struct {
	Auth struct {
		Host        string   `mapstructure:"host"`
		Port        string   `mapstructure:"port"`
	} `mapstructure:"auth"`
	MySQL struct {
		Database   string `mapstructure:"database"`
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
		User string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Automigrate bool `mapstructure:"automigrate"`
		DB          *gorm.DB `mapstructure:"db"`
	} `mapstructure:"mysql"`
	Redis struct {
		URL  string `mapstructure:"url"`
		Pool *redis.Pool
	} `mapstructure:"redis"`
}

// C holds the global configuration.
var C Config
