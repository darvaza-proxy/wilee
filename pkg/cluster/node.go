package cluster

import (
	"darvaza.org/core"

	"github.com/shaj13/raft"
	"github.com/shaj13/raft/transport"
)

// Node ...
type Node struct {
	rn  *raft.Node
	cfg *Config
}

// Start ...
func (n *Node) Start() error {
	opts, err := n.cfg.ExportStartNodeOptions()
	if err != nil {
		return err
	}

	if err := n.rn.Start(opts...); err != nil {
		return core.Wrap(err, "raft.Start")
	}

	return nil
}

// NewNode ...
func NewNode(cfg *Config) (*Node, error) {
	if cfg == nil {
		return nil, core.Wrap(core.ErrInvalid, "Config not provided")
	}

	proto := transport.GRPC

	opts, err := cfg.ExportNewNodeOptions(proto)
	if err != nil {
		return nil, err
	}

	rn := raft.NewNode(&StateMachine{}, proto, opts...)
	n := &Node{
		rn:  rn,
		cfg: cfg,
	}

	for _, fn := range []func(*Node) error{} {
		if err := fn(n); err != nil {
			return nil, err
		}
	}

	return n, nil
}
