package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gravitational/rigging"
	"github.com/gravitational/trace"
)

func install() error {
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
