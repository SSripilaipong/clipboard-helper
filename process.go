package main

import "github.com/SSripilaipong/go-common/rslt"

// Process orchestrates the command execution pipeline.
func Process(cmdReader CommandReader, processor Processor, inputReader InputReader, resultWriter OutputWriter) {
	// TODO: implement pipeline
}

// CommandReader is responsible for obtaining the user's requested operation.
type CommandReader interface {
	ReadCommand() rslt.Of[string]
}

// InputReader fetches the text that will be manipulated.
type InputReader interface {
	ReadInput() rslt.Of[string]
}

// OutputWriter pushes the processed text back to the destination.
type OutputWriter interface {
	WriteOutput(string) error
}

// Processor applies a command to an input and yields the output.
type Processor interface {
	Process(rslt.Of[string], rslt.Of[string]) rslt.Of[string]
}
