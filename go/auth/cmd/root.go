package cmd

import (
	"bytes"
	"io/ioutil"
	log "github.com/cihub/seelog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"auth/config"
)

var cfgFile string // config file
var logXml string // log config file
var version string

var rootCmd = &cobra.Command{
	Use:   "Auth Server",
	Short: "Auth Server project",
	Long:  `Auth Server project`,
	RunE:  run,
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "path to configuration file (optional)")
	// for backwards compatibility
	rootCmd.PersistentFlags().String("mysql-database","auth","database name")
	rootCmd.PersistentFlags().String("mysql-host","192.168.1.6","e.g: 192.168.1.6")
	rootCmd.PersistentFlags().String("mysql-port","3306","e.g: 3306")
	rootCmd.PersistentFlags().String("mysql-user","root","")
	rootCmd.PersistentFlags().String("mysql-password","Radacat2017","")
	rootCmd.PersistentFlags().Bool("mysql-automigrate",false,"true/false")
	rootCmd.PersistentFlags().String("redis-url","redis://root:radacat1234@192.168.1.6:6379/0","e.g. redis://user:password@hostname/0")
	rootCmd.PersistentFlags().String("auth-host","127.0.0.1","")
	rootCmd.PersistentFlags().String("auth-port","7010","")
	// for backwards compatibility


	// for backwards compatibility
	viper.BindPFlag("mysql.database",rootCmd.PersistentFlags().Lookup("mysql-database"))
	viper.BindPFlag("mysql.host",rootCmd.PersistentFlags().Lookup("mysql-host"))
	viper.BindPFlag("mysql.port",rootCmd.PersistentFlags().Lookup("mysql-port"))
	viper.BindPFlag("mysql.user",rootCmd.PersistentFlags().Lookup("mysql-user"))
	viper.BindPFlag("mysql.password",rootCmd.PersistentFlags().Lookup("mysql-password"))
	viper.BindPFlag("mysql.automigrate",rootCmd.PersistentFlags().Lookup("mysql-automigrate"))
	viper.BindPFlag("redis.url",rootCmd.PersistentFlags().Lookup("redis-url"))
	viper.BindPFlag("auth.host",rootCmd.PersistentFlags().Lookup("auth-host"))
	viper.BindPFlag("auth.port",rootCmd.PersistentFlags().Lookup("auth-port"))


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
		viper.SetConfigName("auth")
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
