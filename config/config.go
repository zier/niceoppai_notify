package config

import (
	"os"
	"strings"

	"gopkg.in/urfave/cli.v1"
)

// Config ....
type Config struct {
	CLI       *cli.App
	AppConfig *AppConfig
}

// AppConfig ...
type AppConfig struct {
	Tokens []string
}

// New ...
func New() *Config {
	config := &Config{}
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "tokens, ts",
			Usage: "line notify token, if you have more than one token use for seperate (token_1,token_2)",
		},
	}

	app.Action = func(c *cli.Context) error {
		config.AppConfig = &AppConfig{
			Tokens: strings.Split(c.String("tokens"), ","),
		}
		return nil
	}

	config.CLI = app
	return config
}

// ReadCLIParams ...
func (c *Config) ReadCLIParams() error {
	return c.CLI.Run(os.Args)

}
