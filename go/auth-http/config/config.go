package config

// Config defines the configuration structure.
type Config struct {
	AuthHttp struct {
		Host        string   `mapstructure:"host"`
		Port        string   `mapstructure:"port"`
	} `mapstructure:"auth_http"`
	Auth struct{
		HostName        string   `mapstructure:"hostname"`
	}`mapstructure:"auth"`
}

// C holds the global configuration.
var C Config
