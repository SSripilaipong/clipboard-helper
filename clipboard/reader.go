package clipboard

import (
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/atotto/clipboard"
)

// Reader pulls text from the system clipboard.
type Reader struct{}

// NewReader provides the clipboard-backed InputReader implementation.
func NewReader() *Reader {
	return &Reader{}
}

// ReadInput retrieves the clipboard text.
func (r *Reader) ReadInput() rslt.Of[string] {
	return rslt.New(clipboard.ReadAll())
}
