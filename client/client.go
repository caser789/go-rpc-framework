package client

import "errors"
import "net/rpc"
import "net/rpc/jsonrpc"
import "strconv"

import "github.com/caser789/go-rpc-framework/core"

type Client struct {
	Port uint
    UseHttp bool
    UseJson bool
    client *rpc.Client
}

func (c *Client) Init(name string) (msg string, err error) {
	if c.Port == 0 {
		err = errors.New("client: port must be specified")
		return
	}

	addr := "127.0.0.1:" + strconv.Itoa(int(c.Port))

    if c.UseHttp {
        c.client, err = rpc.DialHTTP("tcp", addr)
    } else if c.UseJson {
        c.client, err = jsonrpc.Dial("tcp", addr)
    } else {
        c.client, err = rpc.Dial("tcp", addr)
    }
	if err != nil {
		return
	}

	return
}

func (c *Client) Close() (err error) {
    if c.client != nil {
        err = c.client.Close()
        return
    }

    return
}

func (c *Client) Execute(name string) (msg string, err error) {
    var request = &core.Request{Name: name}
    var response = new(core.Response)

    err = c.client.Call(core.HandlerName, request, response)
    if err != nil {
        return
    }
}
