package clipboard

import (
	"log"

	"github.com/SSripilaipong/go-common/rslt"
	"github.com/atotto/clipboard"
)

// Writer pushes text back to the system clipboard.
type Writer struct{}

// NewWriter provides the clipboard-backed OutputWriter implementation.
func NewWriter() *Writer {
	return &Writer{}
}

// WriteOutput writes the provided result to the clipboard.
func (w *Writer) WriteOutput(value rslt.Of[string]) {
	if value.IsErr() {
		w.WriteError(value.Error())
		return
	}
	if err := clipboard.WriteAll(value.Value()); err != nil {
		w.WriteError(err)
	}
}

// WriteError logs clipboard errors for now.
func (w *Writer) WriteError(err error) {
	if err != nil {
		log.Printf("clipboard write error: %v", err)
	}
}
