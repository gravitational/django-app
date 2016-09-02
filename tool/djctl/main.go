package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/alecthomas/kingpin"
	"github.com/gravitational/trace"
)

const (
	EnvStolonRPCHost = "STOLON_RPC_SERVICE_HOST"
	EnvStolonRPCPort = "STOLON_RPC_SERVICE_PORT"
	EnvDatabaseName  = "DB_NAME"
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
	var (
		clt    Client
		dbName string
	)
	// install
	cmdInstall := app.Command("install", "install django application")
	cmdInstall.Flag("db-name", "database name").Envar(EnvDatabaseName).StringVar(&dbName)
	cmdInstall.Flag("stolon-rpc-host", "Stolon RPC host").Envar(EnvStolonRPCHost).StringVar(&clt.Host)
	cmdInstall.Flag("stolon-rpc-port", "Stolon RPC port").Envar(EnvStolonRPCPort).StringVar(&clt.Port)
	// uninstall
	cmdUninstall := app.Command("uninstall", "uninstall django application")
	cmdUninstall.Flag("db-name", "database name").Envar(EnvDatabaseName).StringVar(&dbName)
	cmdUninstall.Flag("stolon-rpc-host", "Stolon RPC host").Envar(EnvStolonRPCHost).StringVar(&clt.Host)
	cmdUninstall.Flag("stolon-rpc-port", "Stolon RPC port").Envar(EnvStolonRPCPort).StringVar(&clt.Port)

	cmd, err := app.Parse(os.Args[1:])
	if err != nil {
		return trace.Wrap(err)
	}

	switch cmd {
	case cmdInstall.FullCommand():
		return Install(clt, dbName)
	case cmdUninstall.FullCommand():
		return Uninstall(clt, dbName)
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
