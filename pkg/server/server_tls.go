package server

import (
	"darvaza.org/x/tls/store/basic"
)

func (srv *Server) initTLS() error {
	srv.tls = basic.New()
	return nil
}
