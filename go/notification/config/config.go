package config

import (
	"github.com/jinzhu/gorm"
	"github.com/garyburd/redigo/redis"
	)

var C Config

type Config struct {
	Notification struct{
		Host        string   `mapstructure:"host"`
		Port        string   `mapstructure:"port"`
	}`mapstructure:"notification"`
	MySQL struct {
		Database    string   `mapstructure:"database"`
		Host        string   `mapstructure:"host"`
		Port        string   `mapstructure:"port"`
		User        string   `mapstructure:"user"`
		Password    string   `mapstructure:"password"`
		Automigrate bool     `mapstructure:"automigrate"`
		DB          *gorm.DB `mapstructure:"db"`
	} `mapstructure:"mysql"`
	Redis struct {
		URL  string `mapstructure:"url"`
		Pool *redis.Pool
	} `mapstructure:"redis"`
	MQTT struct {
		Server    string `mapstructure:"server"`
		Username  string `mapstructure:"username"`
		Password  string `mapstructure:"Password"`
		CACert    string `mapstructure:"ca_cert"`
	}`mapstructure:"mqtt"`
	Auth struct {
		Hostname    string `mapstructure:"hostname"`
	}`mapstructure:"auth"`
}

