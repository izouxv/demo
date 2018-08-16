package cmd

import (
	"bytes"
	"io/ioutil"
	log "github.com/cihub/seelog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"notification/config"
	"fmt"
)

var cfgFile string // config file

var logXml string // log config file

var version  string

var rootCmd = &cobra.Command{
	Use:   "Notification Server",
	Short: "Notification Server project ",
	Long:  `Notification Server is an application-server`,
	RunE:  run,
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "path to configuration file (optional)")
	//mysql
	rootCmd.PersistentFlags().String("mysql-database","rule-engine","database name")
	rootCmd.PersistentFlags().String("mysql-host","192.168.1.6","e.g: 127.0.0.1")
	rootCmd.PersistentFlags().String("mysql-port","3306","e.g: 3306")
	rootCmd.PersistentFlags().String("mysql-user","","")
	rootCmd.PersistentFlags().String("mysql-password","","")
	rootCmd.PersistentFlags().Bool(  "mysql-automigrate",false,"true/false")
	rootCmd.PersistentFlags().String("redis-url","","")
	rootCmd.PersistentFlags().String("notification-host","0.0.0.0","")
	rootCmd.PersistentFlags().String("notification-port","7023","")
	rootCmd.PersistentFlags().String("auth-hostname","0.0.0.0:7010","")


	rootCmd.PersistentFlags().String("mqtt-server", "tcp://192.168.1.6:1884", "e.g. scheme://host:port where scheme is tcp, ssl or ws")
	rootCmd.PersistentFlags().String("mqtt-username", "", "")
	rootCmd.PersistentFlags().String("mqtt-password", "", "")
	rootCmd.PersistentFlags().String("mqtt-ca-cert", "", "")

	// for backwards compatibility
	viper.BindPFlag("mysql.database",    rootCmd.PersistentFlags().Lookup("mysql-database"))
	viper.BindPFlag("mysql.host",        rootCmd.PersistentFlags().Lookup("mysql-host"))
	viper.BindPFlag("mysql.port",        rootCmd.PersistentFlags().Lookup("mysql-port"))
	viper.BindPFlag("mysql.user",        rootCmd.PersistentFlags().Lookup("mysql-user"))
	viper.BindPFlag("mysql.password",    rootCmd.PersistentFlags().Lookup("mysql-password"))
	viper.BindPFlag("mysql.automigrate", rootCmd.PersistentFlags().Lookup("mysql-automigrate"))
	viper.BindPFlag("redis.url",         rootCmd.PersistentFlags().Lookup("redis-url"))
	viper.BindPFlag("notification.host", rootCmd.PersistentFlags().Lookup("notification-host"))
	viper.BindPFlag("notification.port", rootCmd.PersistentFlags().Lookup("notification-port"))
	viper.BindPFlag("auth.hostname",     rootCmd.PersistentFlags().Lookup("auth-hostname"))

	viper.BindPFlag("mqtt.server",     rootCmd.PersistentFlags().Lookup("mqtt-server"))
	viper.BindPFlag("mqtt.username",   rootCmd.PersistentFlags().Lookup("mqtt-username"))
	viper.BindPFlag("mqtt.password",   rootCmd.PersistentFlags().Lookup("mqtt-password"))
	viper.BindPFlag("mqtt.ca_cert",    rootCmd.PersistentFlags().Lookup("mqtt-ca-cert"))

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
		viper.SetConfigName("admin")
		viper.AddConfigPath(".")
		viper.AddConfigPath("e://etc/")
		viper.AddConfigPath("/etc/config/")
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
		fmt.Println("err :",err)
		log.Debugf("unmarshal config error,error :%s",err)
	}
}
