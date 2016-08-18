package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gravitational/rigging"
	"github.com/gravitational/trace"
)

func install(dbName string) error {
	log.Infof("creating database %s", dbName)
	if err := createDB(dbName); err != nil {
		log.Errorf("Can't create database '%s'. Err: %v", dbName, err)
	}

	log.Infof("creating django service and replication controller")
	out, err := rigging.FromFile(
		rigging.ActionCreate,
		"/var/lib/gravity/resources/django.yaml")
	if err != nil {
		log.Errorf("%s", string(out))
		return trace.Wrap(err)
	}

	return nil
}

func uninstall() error {
	log.Infof("deleting django service and replication controller")
	out, err := rigging.FromFile(
		rigging.ActionDelete,
		"/var/lib/gravity/resources/django.yaml")
	if err != nil {
		log.Errorf("%s", string(out))
		return trace.Wrap(err)
	}

	return nil
}
