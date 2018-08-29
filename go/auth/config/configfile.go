package config

import (
	"html/template"
	"os"
	"github.com/spf13/cobra"
	"github.com/pkg/errors"
)

// when updating this template, don't forget to update config.md!
const configTemplate = `
[auth]
	host={{ .Auth.Host }}
	port={{ .Auth.Port }}
[mysql]
	database={{ .MySQL.Database }}
	user={{ .MySQL.User  }}
	password={{ .MySQL.Password }}
	host={{ .MySQL.Host }}
	port={{ .MySQL.Port }}
    automigrate= {{ .MySQL.Automigrate }}
[redis]
	url= {{ .Redis.URL }}
`

const LogConfig = `
<seelog>
    <outputs formatid="main">
        <filter levels="debug,info,error">
            <console />
        </filter>
        <filter levels="debug">
            <rollingfile formatid="main" type="date" filename="/var/log/golang/auth_debug.log" datepattern="02.01.2006" maxrolls="7"/>
        </filter>
        <filter levels="info">
            <rollingfile formatid="main" type="date" filename="/var/log/golang/auth_info.log" datepattern="02.01.2006" maxrolls="7"/>
        </filter>
        <filter levels="error">
            <rollingfile formatid="main" type="date" filename="/var/log/golang/auth_error.log" datepattern="02.01.2006" maxrolls="7" />
        </filter>
    </outputs>
    <formats>
        <format format="%Date %Time|%LEV|%File:%Line|%Msg%n" id32="main"/>
    </formats>
</seelog>
`
var ConfigCmd = &cobra.Command{
	Use:   "configfile",
	Short: "Print the Auth Server configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		t := template.Must(template.New("config").Parse(configTemplate))
		err := t.Execute(os.Stdout, C)
		if err != nil {
			return errors.Wrap(err, "execute config template error")
		}
		return nil
	},
}
