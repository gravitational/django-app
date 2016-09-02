package main

import (
	"fmt"
	"net"
	"net/rpc"

	log "github.com/Sirupsen/logrus"
	"github.com/gravitational/rigging"
	"github.com/gravitational/trace"
)

type Client struct {
	Host, Port string
}

func (c *Client) Get() (*rpc.Client, error) {
	client, err := rpc.DialHTTP("tcp", c.connectionString())
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return client, nil
}

func (c *Client) connectionString() string {
	return net.JoinHostPort(c.Host, c.Port)
}

func Install(c Client, dbName string) error {
	log.Infof("Creating database '%s'", dbName)
	client, err := c.Get()
	if err != nil {
		return trace.Wrap(err, fmt.Sprintf("Dialing to stolon's RPC at '%s' failed", c.connectionString()))
	}

	var reply string
	command := "DatabaseOperation.Create"
	err = client.Call(command, dbName, &reply)
	log.Infof("Execute RPC command '%s' on stolon's RPC at '%s'", command, c.connectionString())
	if err != nil {
		return trace.Wrap(err, fmt.Sprintf("Can't create database '%s'", dbName))
	}
	log.Infof("Reply: %s", reply)

	log.Infof("Creating django service and replication controller")
	out, err := rigging.FromFile(rigging.ActionCreate, "/var/lib/gravity/resources/django.yaml")
	if err != nil {
		return trace.Wrap(err, fmt.Sprintf("cmd output: %s", string(out)))
	}

	return nil
}

func Uninstall(c Client, dbName string) error {
	log.Infof("Deleting django service and replication controller")
	out, err := rigging.FromFile(
		rigging.ActionDelete,
		"/var/lib/gravity/resources/django.yaml")
	if err != nil {
		return trace.Wrap(err, fmt.Sprintf("cmd output: %s", string(out)))
	}

	log.Infof("Deleting database '%s'", dbName)
	client, err := c.Get()
	if err != nil {
		return trace.Wrap(err, fmt.Sprintf("Dialing to stolon's RPC at '%s' failed", c.connectionString()))
	}

	var reply string
	command := "DatabaseOperation.Delete"
	err = client.Call(command, dbName, &reply)
	log.Infof("Execute RPC command '%s' on stolon's RPC at '%s'", command, c.connectionString())
	if err != nil {
		return trace.Wrap(err, fmt.Sprintf("Can't delete database '%s'", dbName))
	}

	log.Infof("Reply: %s", reply)
	return nil
}
