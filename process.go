package main

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	"github.com/SSripilaipong/clipboard-helper/core"
	"github.com/SSripilaipong/clipboard-helper/util/rsltutil"
)

// Process orchestrates the command execution pipeline.
func Process(cmdReader CommandReader, processor Processor, inputReader InputReader, resultWriter OutputWriter) {
	inputs, err := rsltutil.ResultOf2(cmdReader.ReadCommand(), inputReader.ReadInput()).Return()
	if err != nil {
		resultWriter.WriteError(err)
		return
	}
	resultWriter.WriteOutput(tuple.Fn2(processor.Process)(inputs))
}

// CommandReader is responsible for obtaining the user's requested operation.
type CommandReader interface {
	ReadCommand() rslt.Of[core.Command]
}

// InputReader fetches the text that will be manipulated.
type InputReader interface {
	ReadInput() rslt.Of[string]
}

// OutputWriter pushes the processed text back to the destination.
type OutputWriter interface {
	WriteOutput(rslt.Of[string])
	WriteError(error)
}

// Processor applies a command to an input and yields the output.
type Processor interface {
	Process(core.Command, string) rslt.Of[string]
}
