package config

import (
	"html/template"
	"os"
	"github.com/spf13/cobra"
	"github.com/pkg/errors"
)

// when updating this template, don't forget to update config.md!
const configTemplate = `
[auth_http]
	host={{ .AuthHttp.Host }}
	port={{ .AuthHttp.Port }}
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
            <rollingfile formatid="main" type="date" filename="/var/log/golang/auth-http_debug.log" datepattern="02.01.2006" maxrolls="7"/>
        </filter>
        <filter levels="info">
            <rollingfile formatid="main" type="date" filename="/var/log/golang/auth-http_info.log" datepattern="02.01.2006" maxrolls="7"/>
        </filter>
        <filter levels="error">
            <rollingfile formatid="main" type="date" filename="/var/log/golang/auth-http_error.log" datepattern="02.01.2006" maxrolls="7" />
        </filter>
    </outputs>
    <formats>
        <format id="main" format="%Date %Time|%LEV|%File:%Line|%Msg%n"/>
    </formats>
</seelog>
`
var ConfigCmd = &cobra.Command{
	Use:   "configfile",
	Short: "Print the Auth-http Server configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		t := template.Must(template.New("config").Parse(configTemplate))
		err := t.Execute(os.Stdout, C)
		if err != nil {
			return errors.Wrap(err, "execute config template error")
		}
		return nil
	},
}
