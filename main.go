package main

import (
	"github.com/SSripilaipong/clipboard-helper/clipboard"
	"github.com/SSripilaipong/clipboard-helper/processor"
	"github.com/SSripilaipong/clipboard-helper/stdio"
)

func main() {
	cmdReader := stdio.NewCommandReader()
	cmdProcessor := processor.New()
	inputReader := clipboard.NewReader()
	outputWriter := clipboard.NewWriter()
	Process(cmdReader, cmdProcessor, inputReader, outputWriter)
}
