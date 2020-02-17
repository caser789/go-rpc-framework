package client

import "errors"
import "net/rpc"
import "strconv"
import "github.com/caser789/go-rpc-framework/core"

type Client struct {
	Port uint
    UseHttp bool
}

func (c *Client) Execute(name string) (msg string, err error) {
	var client *rpc.Client
	var request = &core.Request{Name: name}
	var response = new(core.Response)

	if c.Port == 0 {
		err = errors.New("client: port must be specified")
		return
	}

	addr := "127.0.0.1:" + strconv.Itoa(int(c.Port))

    if c.UseHttp {
        client, err = rpc.DialHTTP("tcp", addr)
    } else {
        client, err = rpc.Dial("tcp", addr)
    }
	if err != nil {
		return
	}

	err = client.Call("Handler.SayHello", request, response)
	if err != nil {
		return
	}

	msg = response.Message
	return
}
