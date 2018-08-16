package config

import (
	"html/template"
	"os"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"fmt"
)

const ConfigTemplate = `
	[notification]
	host={{ .Notification.Host }}
	port={{ .Notification.Port }}
	[mysql]
	database={{ .MySQL.Database }}
	user={{ .MySQL.User  }}
	password={{ .MySQL.Password }}
	host={{ .MySQL.Host }}
	port={{ .MySQL.Port }}
    automigrate= {{ .MySQL.Automigrate }}
	[redis]
	url= {{ .Redis.URL }}
	[mqtt]
     server={{ .MQTT.Server }}
     username={{ .MQTT.Username }}
     password={{ .MQTT.Password }}
     ca_cert={{ .MQTT.CACert }}
    [auth]
    hostname = {{ .Auth.Hostname}}
`

const LogConfig = `
<seelog>
    <outputs formatid="main">
        <filter levels="debug,info,error">
            <console />
        </filter>
        <filter levels="debug">
            <rollingfile formatid="main" type="date" filename="/var/log/golang/notification_debug.log" datepattern="02.01.2006" maxrolls="7"/>
        </filter>
        <filter levels="info">
            <rollingfile formatid="main" type="date" filename="/var/log/golang/notification_info.log" datepattern="02.01.2006" maxrolls="7"/>
        </filter>
        <filter levels="error">
            <rollingfile formatid="main" type="date" filename="/var/log/golang/notification_error.log" datepattern="02.01.2006" maxrolls="7" />
        </filter>
    </outputs>
    <formats>
        <format format="%Date %Time|%LEV|%File:%Line|%Msg%n" id32="main"/>
    </formats>
</seelog>
`

var ConfigCmd = &cobra.Command{
	Use:   "configfile",
	Short: "Print the notification Server configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		t := template.Must(template.New("config").Parse(ConfigTemplate))
		err := t.Execute(os.Stdout,C)
		fmt.Println("ConfigTemplate....")
		if err != nil {
			return errors.Wrap(err, "execute config template error")
		}
		return nil
	},
}
