package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/gravitational/trace"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type application struct {
	*kingpin.Application
}

func new() *application {
	app := kingpin.New("controller", "command line client")

	var debug bool
	app.Flag("debug", "Enable verbose logging to stderr").
		Short('d').
		BoolVar(&debug)
	if debug {
		InitLoggerDebug()
	} else {
		InitLoggerCLI()
	}

	return &application{app}
}

func (app *application) run() error {
	cmdInstall := app.Command("install", "install cluster")
	cmdUninstall := app.Command("uninstall", "uninstall cluster")

	cmd, err := app.Parse(os.Args[1:])
	if err != nil {
		return trace.Wrap(err)
	}

	switch cmd {
	case cmdInstall.FullCommand():
		return install()
	case cmdUninstall.FullCommand():
		return uninstall()
	}

	return nil
}

func main() {
	app := new()
	if err := app.run(); err != nil {
		log.Error(trace.DebugReport(err))
		os.Exit(1)
	}
}
