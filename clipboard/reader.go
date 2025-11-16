package clipboard

import (
	"errors"

	"github.com/SSripilaipong/go-common/rslt"
)

// Reader pulls text from the system clipboard.
type Reader struct{}

// NewReader provides the clipboard-backed InputReader implementation.
func NewReader() *Reader {
	return &Reader{}
}

// ReadInput retrieves the clipboard text.
func (r *Reader) ReadInput() rslt.Of[string] {
	return rslt.Error[string](errors.New("clipboard reader not implemented"))
}
