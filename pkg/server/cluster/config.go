package cluster

import (
	"context"

	"darvaza.org/core"
	"darvaza.org/sidecar/pkg/logger"
	"darvaza.org/x/config"
)

// Config describes the RAFT cluster's operation
type Config struct {
	Context context.Context
	Logger  Logger
}

// SetDefaults fills any gap in the config.
func (cfg *Config) SetDefaults() error {
	if cfg == nil {
		return core.ErrNilReceiver
	}

	if cfg.Context == nil {
		cfg.Context = context.Background()
	}

	if cfg.Logger.Logger == nil {
		cfg.Logger.Logger = logger.New(nil)
	}

	return config.Set(cfg)
}

// Validate checks the [Config] is valid.
func (cfg *Config) Validate() error {
	var errs core.CompoundError

	if cfg == nil {
		return core.ErrNilReceiver
	}

	if cfg.Context == nil {
		errs.Append(core.ErrInvalid, "missing context")
	}

	if cfg.Logger.Logger == nil {
		errs.Append(core.ErrInvalid, "missing logger")
	}

	return errs.AsError()
}
