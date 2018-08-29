package cmd

import (
	"bytes"
	"io/ioutil"
	log "github.com/cihub/seelog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"auth-http/config"
	"fmt"
)

var cfgFile string // config file
var logXml string // log config file
var version string

var rootCmd = &cobra.Command{
	Use:   "Auth-http Server",
	Short: "Auth-http Server project",
	Long:  `Auth-http Server project`,
	RunE:  run,
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "path to configuration file (optional)")
	// for backwards compatibility
	fmt.Println("cfgFile:",cfgFile)
	rootCmd.PersistentFlags().String("auth_http-host","127.0.0.1","")
	rootCmd.PersistentFlags().String("auth_http-port","7021","")
	rootCmd.PersistentFlags().String("auth-hostname","127.0.0.1:7010","")
	// for backwards compatibility
	viper.BindPFlag("auth_http.host",rootCmd.PersistentFlags().Lookup("auth_http-host"))
	viper.BindPFlag("auth_http.port",rootCmd.PersistentFlags().Lookup("auth_http-port"))
	viper.BindPFlag("auth.hostname",rootCmd.PersistentFlags().Lookup("auth-hostname"))
	// default values
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(config.ConfigCmd)
}

// Execute executes the root command.
func Execute(v string) {
	version = v
	if err := rootCmd.Execute(); err != nil {
		log.Debug(err)
	}
}

func initConfig() {
	if cfgFile != "" {
		b, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			log.Debugf("error loading config file config :%s,err:%s",cfgFile,err)
		}
		viper.SetConfigType("toml")
		if err = viper.ReadConfig(bytes.NewBuffer(b)); err != nil {
			log.Debugf("error loading config file config :%s,err:%s",cfgFile,err)
		}
	} else {
		viper.SetConfigName("auth_http")
		viper.AddConfigPath(".")
		viper.AddConfigPath("e://etc/")
		viper.AddConfigPath("/etc/config/")
		viper.AddConfigPath("./config/")
		if err := viper.ReadInConfig(); err != nil {
			switch err.(type) {
			case viper.ConfigFileNotFoundError:
				log.Debug("Deprecation warning! no configuration file found, falling back on environment variables. Update your configuration")
			default:
				log.Debugf("read configuration file error,error :%s",err)
			}
		}
	}

	if err := viper.Unmarshal(&config.C); err != nil {
		log.Debugf("unmarshal config error,error :%s",err)
	}
}
