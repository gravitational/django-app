package main

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/gravitational/trace"
)

type ServiceMethod string

const (
	OperationDBCreate ServiceMethod = "DatabaseOperation.Create"
	OperationDBDelete ServiceMethod = "DatabaseOperation.Delete"
	OperationDBBackup ServiceMethod = "DatabaseOperation.Backup"
)

type Client struct {
	*rpc.Client
}

type OperationDBBackupArgs struct {
	Name, Path string
}

func NewRPCClient(host, port string) (*Client, error) {
	endpoint := net.JoinHostPort(host, port)
	client, err := rpc.DialHTTP("tcp", endpoint)
	if err != nil {
		return nil, trace.Wrap(err, fmt.Sprintf("Dialing to stolon's RPC at '%s' failed", endpoint))
	}

	return &Client{client}, nil
}

func (c *Client) Execute(cmd ServiceMethod, args interface{}) (string, error) {
	var reply string
	err := c.Call(string(cmd), args, &reply)
	if err != nil {
		return "", trace.Wrap(err, fmt.Sprintf("Procedure '%s' failed", cmd))
	}

	return reply, nil
}
