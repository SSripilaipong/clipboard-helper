package stdio

import (
	"errors"

	"github.com/SSripilaipong/clipboard-helper/core"
	"github.com/SSripilaipong/go-common/rslt"
)

// CommandReader obtains text commands from stdin.
type CommandReader struct{}

// NewCommandReader wires the stdin-backed implementation.
func NewCommandReader() *CommandReader {
	return &CommandReader{}
}

// ReadCommand fetches the next command from the user.
func (s *CommandReader) ReadCommand() rslt.Of[core.Command] {
	return rslt.Error[core.Command](errors.New("stdin command reader not implemented"))
}
