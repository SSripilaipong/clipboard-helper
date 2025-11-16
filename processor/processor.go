package processor

import (
	"errors"

	"github.com/SSripilaipong/clipboard-helper/core"
	"github.com/SSripilaipong/go-common/rslt"
)

// Processor executes transformations based on commands.
type Processor struct{}

// New wires up the clipboard-helper processor implementation.
func New() *Processor {
	return &Processor{}
}

// Process applies the given command to the provided input string.
func (p *Processor) Process(command core.Command, input string) rslt.Of[string] {
	return rslt.Error[string](errors.New("processor not implemented"))
}
