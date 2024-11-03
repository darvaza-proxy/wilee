package server

import (
	"context"
	"net/netip"

	"darvaza.org/core"
	"darvaza.org/slog"
	"darvaza.org/x/config"

	"darvaza.org/sidecar/pkg/logger"
	"darvaza.org/sidecar/pkg/sidecar"
	"darvaza.org/wilee/pkg/server/cluster"
)

// Config describes how the [Server] will operate
type Config struct {
	Context context.Context `yaml:"-" toml:"-" json:"-"`
	Logger  slog.Logger     `yaml:"-" toml:"-" json:"-"`

	// RAFT
	Cluster cluster.Config

	// Supervision provides details for graceful upgrades and restarts
	Supervision sidecar.SupervisionConfig `yaml:"supervision,omitempty" toml:",omitempty" json:",omitempty"`

	// Listen provides details on the addresses and ports to handle
	Listen ListenConfig `yaml:",omitempty" toml:",omitempty" json:",omitempty"`
}

// ListenConfig describes the addresses and ports this [Server] will listen
type ListenConfig struct {
	Interfaces []string     `yaml:",omitempty" toml:",omitempty" json:",omitempty"`
	Addresses  []netip.Addr `yaml:",omitempty"  toml:",omitempty" json:",omitempty"`

	HTTP  uint16 `yaml:"http_port"  default:"80"`
	HTTPS uint16 `yaml:"https_port" default:"443"`
}

// SetDefaults fills any gap in the [Config]
func (cfg *Config) SetDefaults() error {
	if cfg.Context == nil {
		cfg.Context = context.Background()
	}

	if cfg.Logger == nil {
		cfg.Logger = logger.New(nil)
	}

	if cfg.Cluster.Context == nil {
		cfg.Cluster.Context = cfg.Context
	}

	if cfg.Cluster.Logger.Logger == nil {
		cfg.Cluster.Logger.Logger = cfg.Logger
	}

	return config.Set(cfg)
}

// Validate checks if the [Config] is good to use.
func (cfg *Config) Validate() error {
	var errs core.CompoundError

	if cfg == nil {
		return core.ErrNilReceiver
	}

	if cfg.Context == nil {
		errs.Append(core.ErrInvalid, "context missing")
	}

	if cfg.Logger == nil {
		errs.Append(core.ErrInvalid, "logger missing")
	}

	if err := cfg.Cluster.Validate(); err != nil {
		errs.Append(err, "cluster")
	}

	return errs.AsError()
}
