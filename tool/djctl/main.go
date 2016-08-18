package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/alecthomas/kingpin"
	"github.com/gravitational/trace"
)

const (
	EnvDatabaseName = "DB_NAME"
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
	var dbName string
	// install
	cmdInstall := app.Command("install", "install django application")
	cmdInstall.Flag("db-name", "database name").Envar(EnvDatabaseName).StringVar(&dbName)
	// uninstall
	cmdUninstall := app.Command("uninstall", "uninstall django application")

	cmd, err := app.Parse(os.Args[1:])
	if err != nil {
		return trace.Wrap(err)
	}

	switch cmd {
	case cmdInstall.FullCommand():
		return install(dbName)
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
