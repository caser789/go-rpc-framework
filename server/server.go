package server

import "errors"
import "net"
import "net/rpc"
import "net/http"
import "strconv"
import "github.com/caser789/go-rpc-framework/core"

type Server struct {
	Port     uint
	listener net.Listener
    UseHttp bool
}

func (s *Server) Stop() (err error) {
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

	rpc.Register(new(core.Handler))

	s.listener, err = net.Listen("tcp", ":"+strconv.Itoa(int(s.Port)))
	if err != nil {
		return
	}

    if s.UseHttp {
        rpc.HandleHTTP()
        http.Serve(s.listener, nil)
    } else {
        rpc.Accept(s.listener)
    }
	return
}
