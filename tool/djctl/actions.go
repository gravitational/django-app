package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/gravitational/rigging"
	"github.com/gravitational/trace"
)

func Install(c *Client, dbName string) error {
	log.Infof("Creating database '%s', call procedure '%s'", dbName, OperationDBCreate)
	rpcReply, err := c.Execute(OperationDBCreate, dbName)
	if err != nil {
		return trace.Wrap(err)
	}
	log.Infof("Reply: %s", rpcReply)

	log.Infof("Creating Service and Replication Controller")
	kubeReply, err := rigging.FromFile(rigging.ActionCreate, "/var/lib/gravity/resources/django.yaml")
	if err != nil {
		return trace.Wrap(err, fmt.Sprintf("Reply: %s", string(kubeReply)))
	}
	log.Info("Done.")

	return nil
}

func Update(c *Client, dbName, dbBackupPath string) error {
	log.Infof("Backup database '%s' to '%s', call procedure '%s'", dbName, dbBackupPath, OperationDBBackup)
	rpcReply, err := c.Execute(OperationDBBackup, OperationDBBackupArgs{Name: dbName, Path: dbBackupPath})
	if err != nil {
		return trace.Wrap(err)
	}
	log.Infof("Reply: %s", rpcReply)

	return nil
}

func Uninstall(c *Client, dbName string) error {
	log.Infof("Deleting django service and replication controller")
	kubeReply, err := rigging.FromFile(rigging.ActionDelete, "/var/lib/gravity/resources/django.yaml")
	if err != nil {
		return trace.Wrap(err, fmt.Sprintf("Reply: %s", string(kubeReply)))
	}

	log.Infof("Deleting database '%s', call procedure '%s'", dbName, OperationDBDelete)
	rpcReply, err := c.Execute(OperationDBDelete, dbName)
	if err != nil {
		return trace.Wrap(err)
	}
	log.Infof("Reply: %s", rpcReply)

	log.Info("Done.")

	return nil
}
