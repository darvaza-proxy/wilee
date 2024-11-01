package cluster

import (
	"errors"

	"github.com/shaj13/raft"
	"github.com/shaj13/raft/transport"

	"darvaza.org/core"
)

// ExportNewNodeOptions returns the options to be passed to [raft.NewNode]
// according to the [Config].
func (cfg *Config) ExportNewNodeOptions(_ transport.Proto) ([]raft.Option, error) {
	var out []raft.Option

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	out = append(out, raft.WithContext(cfg.Context))
	out = append(out, raft.WithLogger(&cfg.Logger))

	return out, nil
}

// ExportStartNodeOptions returns the options to be passed to [raft.Node#Start]
// according to the [Config].
func (cfg *Config) ExportStartNodeOptions() ([]raft.StartOption, error) {
	var out []raft.StartOption

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return out, nil
}

// NewNode creates a [raft.Node] according to the [Config].
func (cfg *Config) NewNode(sm raft.StateMachine, proto transport.Proto) (*raft.Node, error) {
	if sm == nil {
		return nil, core.Wrap(core.ErrInvalid, "raft state machine missing")
	}

	opts, err := cfg.ExportNewNodeOptions(proto)
	if err != nil {
		return nil, err
	}

	n := raft.NewNode(sm, proto, opts...)
	if n == nil {
		return nil, errors.New("raft failed to create a Node")
	}
	return n, nil
}

// StartNode starts a [raft.Node] according to the [Config].
func (cfg *Config) StartNode(n *raft.Node) error {
	opts, err := cfg.ExportStartNodeOptions()
	if err != nil {
		return err
	}

	return n.Start(opts...)
}
