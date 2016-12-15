package main

import (
	"fmt"
	sakura "github.com/yamamoto-febc/sakura-iot-go"
	"github.com/yamamoto-febc/sakura-iot-go/version"
	"gopkg.in/urfave/cli.v2"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type option struct {
	HostName string
	Path     string
	Port     int
	Secret   string
	Debug    bool
}

func (o *option) validate() []error {
	ret := []error{}

	if o.Path == "" {
		ret = append(ret, fmt.Errorf("%s is required", "--path"))
	}

	if o.Port < 1 || 65535 < o.Port {
		ret = append(ret, fmt.Errorf("%s is neet between 1 to 65535", "--port"))
	}

	return ret
}

func newOption() *option {
	return &option{}
}

var (
	appName              = "sakura-iot-go:echo_server"
	appUsage             = "Echo all posts from Sakura IoT platform"
	appCopyright         = "Copyright (C) 2016 Kazumichi Yamamoto."
	applHelpTextTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.HelpName}} [options]

OPTIONS:
   {{range .OptionFlags}}{{.}}
   {{end}}
VERSION:
   {{.Version}}

{{.Copyright}}
`
)

func main() {

	cli.AppHelpTemplate = applHelpTextTemplate
	app := &cli.App{}
	option := newOption()

	app.Name = appName
	app.Usage = appUsage
	app.HelpName = appName
	app.Copyright = appCopyright

	app.Flags = cliFlags(option)
	app.Action = cliCommand(option)
	app.Version = version.FullVersion()

	originalHelpPrinter := cli.HelpPrinter
	cli.HelpPrinter = func(w io.Writer, templ string, d interface{}) {
		app := d.(*cli.App)
		data := newHelpData(app)
		originalHelpPrinter(w, templ, data)
	}

	app.Run(os.Args)
}

type helpData struct {
	*cli.App
	RequiredFlags []cli.Flag
	OptionFlags   []cli.Flag
}

func newHelpData(app *cli.App) interface{} {
	data := &helpData{App: app}

	for _, f := range app.VisibleFlags() {
		data.OptionFlags = append(data.OptionFlags, f)
	}

	return data
}

func cliFlags(option *option) []cli.Flag {

	return []cli.Flag{
		&cli.StringFlag{
			Name:        "path",
			EnvVars:     []string{"SAKURA_IOT_ECHO_PATH"},
			DefaultText: "/",
			Value:       "/",
			Destination: &option.Path,
			Usage:       "Request receive path",
		},
		&cli.StringFlag{
			Name:        "hostname",
			EnvVars:     []string{"SAKURA_IOT_ECHO_HOSTNAME"},
			DefaultText: "",
			Destination: &option.HostName,
			Usage:       "Listen hostname or ipaddress",
		},
		&cli.IntFlag{
			Name:        "port",
			EnvVars:     []string{"SAKURA_IOT_ECHO_PORT"},
			DefaultText: "8080",
			Value:       8080,
			Destination: &option.Port,
			Usage:       "Listen port",
		},
		&cli.StringFlag{
			Name:        "secret",
			EnvVars:     []string{"SAKURA_IOT_ECHO_SECRET"},
			DefaultText: "",
			Destination: &option.Secret,
			Usage:       "secret",
		},
		&cli.BoolFlag{
			Name:        "debug",
			Usage:       "Flag of enable DEBUG log",
			EnvVars:     []string{"SAKURA_IOT_ECHO_DEBUG"},
			Destination: &option.Debug,
			Value:       false,
		},
	}

}

func cliCommand(option *option) func(c *cli.Context) error {
	return func(c *cli.Context) error {

		errors := option.validate()
		if len(errors) != 0 {
			return flattenErrors(errors)
		}

		log.SetFlags(log.Ldate | log.Ltime)
		log.SetOutput(os.Stdout)
		out := log.Printf

		handler := &sakura.WebhookHandler{
			Secret: option.Secret,
			ConnectedFunc: func(p sakura.Payload) {
				out("[INFO] Connected module message received:\n%#v", p)
			},
			HandleFunc: func(p sakura.Payload) {
				out("[INFO] Outgoing Webhook received:\n%#v", p)
			},
			Debug: option.Debug,
		}

		addr := fmt.Sprintf("%s:%d", option.HostName, option.Port)

		out("[INFO] start ListenAndServe. addr:[%s] path:[%s] secret:[%s]\n", addr, option.Path, option.Secret)
		http.Handle(option.Path, handler)
		return http.ListenAndServe(addr, nil)

	}
}

func flattenErrors(errors []error) error {
	var list = make([]string, 0)
	for _, str := range errors {
		list = append(list, str.Error())
	}
	return fmt.Errorf(strings.Join(list, "\n"))
}

func isExistsFlag(source []string, target cli.Flag) bool {
	for _, s := range source {
		if s == target.Names()[0] {
			return true
		}
	}
	return false
}
