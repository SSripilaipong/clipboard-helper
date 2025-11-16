package clipboard

import "errors"

// Writer pushes text back to the system clipboard.
type Writer struct{}

// NewWriter provides the clipboard-backed OutputWriter implementation.
func NewWriter() *Writer {
	return &Writer{}
}

// WriteOutput writes the provided text to the clipboard.
func (w *Writer) WriteOutput(value string) error {
	return errors.New("clipboard writer not implemented")
}
