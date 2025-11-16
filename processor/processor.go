package processor

import (
	"errors"

	"github.com/SSripilaipong/go-common/rslt"
)

// Processor executes transformations based on commands.
type Processor struct{}

// New wires up the clipboard-helper processor implementation.
func New() *Processor {
	return &Processor{}
}

// Process applies the given command to the provided input string.
func (p *Processor) Process(command rslt.Of[string], input rslt.Of[string]) rslt.Of[string] {
	return rslt.Error[string](errors.New("processor not implemented"))
}
