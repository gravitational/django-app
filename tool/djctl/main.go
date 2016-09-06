package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/alecthomas/kingpin"
	"github.com/gravitational/trace"
)

const (
	EnvStolonRPCHost      = "STOLON_RPC_SERVICE_HOST"
	EnvStolonRPCPort      = "STOLON_RPC_SERVICE_PORT"
	EnvDatabaseName       = "DB_NAME"
	EnvDatabaseBackupPath = "DB_BACKUP_PATH"
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
	cmdApp := app.Command("app", "operations with django application")
	rpcHost := cmdApp.Flag("stolon-rpc-host", "Stolon RPC host").Envar(EnvStolonRPCHost).Required().String()
	rpcPort := cmdApp.Flag("stolon-rpc-port", "Stolon RPC port").Envar(EnvStolonRPCPort).Required().String()
	dbName := cmdApp.Flag("db-name", "database name").Envar(EnvDatabaseName).Required().String()

	cmdAppInstall := cmdApp.Command("install", "install django application")
	cmdAppUninstall := cmdApp.Command("uninstall", "uninstall django application")
	cmdAppUpdate := cmdApp.Command("update", "update django application")
	dbBackupPath := cmdAppUpdate.Flag("db-backup-path", "Path where database backup will be stored").Envar(EnvDatabaseBackupPath).Required().String()

	cmd, err := app.Parse(os.Args[1:])
	if err != nil {
		return trace.Wrap(err)
	}

	clt, err := NewRPCClient(*rpcHost, *rpcPort)
	if err != nil {
		return trace.Wrap(err)
	}

	switch cmd {
	case cmdAppInstall.FullCommand():
		return Install(clt, *dbName)
	case cmdAppUpdate.FullCommand():
		return Update(clt, *dbName, *dbBackupPath)
	case cmdAppUninstall.FullCommand():
		return Uninstall(clt, *dbName)
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
