package server

import "errors"
import "net"
import "net/rpc"
import "net/rpc/jsonrpc"
import "net/http"
import "strconv"
import "time"

import "github.com/caser789/go-rpc-framework/core"

type Server struct {
	Port     uint
	listener net.Listener
    UseHttp bool
    UseJson bool
    Sleep time.Duration
}

func (s *Server) Close() (err error) {
	if s.listener != nil {
		err = s.listener.Close()
	}

	return
}

func (s *Server) Start() (err error) {
	if s.Port <= 0 {
		err = errors.New("port must be specified")
		return
	}

	rpc.Register(&core.Handler{
        Sleep: s.Sleep,
    })

	s.listener, err = net.Listen("tcp", ":"+strconv.Itoa(int(s.Port)))
	if err != nil {
		return
	}

    if s.UseHttp {
        rpc.HandleHTTP()
        http.Serve(s.listener, nil)
    } else if s.UseJson {
        var conn net.Conn

        for {
            conn, err = s.listener.Accept()
            if err != nil {
                return
            }
            jsonrpc.ServeConn(conn)
        }
    } else {
        rpc.Accept(s.listener)
    }
	return
}
