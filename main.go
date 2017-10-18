package main

//go:generate esc -o static.go -pkg main -prefix ui ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/ghetzel/cli"
	"github.com/ghetzel/diecast"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger(`main`)

type OnQuitFunc func() // {}
var OnQuit OnQuitFunc

func main() {
	app := cli.NewApp()
	app.Name = `eval-server`
	app.Usage = `A server that serves things.`
	app.Version = `0.0.1`
	app.EnableBashCompletion = false

	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:   `log-level, L`,
			Usage:  `Level of log output verbosity`,
			EnvVar: `LOGLEVEL`,
		},
		cli.StringFlag{
			Name:  `ui-dir`,
			Usage: `The path to the UI directory`,
			Value: `embedded`,
		},
		cli.StringFlag{
			Name:  `address, a`,
			Usage: `The address the server should listen on`,
			Value: `0.0.0.0:8086`,
		},
	}

	app.Before = func(c *cli.Context) error {
		var addlInfo string
		levels := append([]string{
			`debug`,
		}, c.StringSlice(`log-level`)...)

		for _, levelspec := range levels {
			var levelName string
			var moduleName string

			if parts := strings.SplitN(levelspec, `:`, 2); len(parts) == 1 {
				levelName = parts[0]
			} else {
				moduleName = parts[0]
				levelName = parts[1]
			}

			if level, err := logging.LogLevel(levelName); err == nil {
				if level == logging.DEBUG {
					addlInfo = `%{module}: `
				}

				logging.SetLevel(level, moduleName)
			} else {
				return err
			}
		}

		logging.SetFormatter(logging.MustStringFormatter(
			fmt.Sprintf("%%{color}%%{level:.4s}%%{color:reset}[%%{id:04d}] %s%%{message}", addlInfo),
		))

		log.Infof("Starting %s %s", c.App.Name, c.App.Version)

		return nil
	}

	app.Action = func(c *cli.Context) {
		server := diecast.NewServer(c.String(`ui-dir`), `*.html`)
		address := c.String(`address`)

		if strings.HasPrefix(address, `:`) {
			server.BindingPrefix = fmt.Sprintf("http://localhost%s", address)
		} else {
			server.BindingPrefix = fmt.Sprintf("http://%s", address)
		}

		if c.String(`ui-dir`) == `embedded` {
			server.SetFileSystem(FS(false))
		}

		if err := server.Initialize(); err == nil {
			server.ListenAndServe(address)
		} else {
			log.Fatalf("Failed to start server: %v", err)
		}
	}

	app.Run(os.Args)
}
