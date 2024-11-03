// Package server implements the Wilee server
package server

import (
	"darvaza.org/x/tls"
)

// Server is a Wilee server
type Server struct {
	cfg Config

	tls tls.Store
}

func (srv *Server) init() error {
	for _, fn := range []func() error{
		srv.cfg.SetDefaults,
		srv.initTLS,
	} {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}

// New creates a [Server] using the given [Config]
func New(cfg *Config) (*Server, error) {
	if cfg == nil {
		cfg = new(Config)
	}

	srv := &Server{
		cfg: *cfg,
	}

	if err := srv.init(); err != nil {
		return nil, err
	}

	return srv, nil
}
