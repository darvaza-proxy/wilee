package cluster

import (
	"io"

	"github.com/shaj13/raft"

	"darvaza.org/core"
)

var _ raft.StateMachine = (*StateMachine)(nil)

// StateMachine implements the RAFT cluster state machine
type StateMachine struct{}

// Apply implements the [raft.StateMachine] interface
func (*StateMachine) Apply([]byte) {}

// Restore implements the [raft.StateMachine] interface
func (*StateMachine) Restore(io.ReadCloser) error {
	return core.ErrTODO
}

// Snapshot implements the [raft.StateMachine] interface
func (*StateMachine) Snapshot() (io.ReadCloser, error) {
	return nil, core.ErrTODO
}
