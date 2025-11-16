package stdio

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/SSripilaipong/clipboard-helper/core"
	"github.com/SSripilaipong/go-common/rslt"
)

// CommandReader obtains text commands from stdin.
type CommandReader struct {
	reader *bufio.Reader
}

// NewCommandReader wires the stdin-backed implementation.
func NewCommandReader() *CommandReader {
	return &CommandReader{
		reader: bufio.NewReader(os.Stdin),
	}
}

// ReadCommand fetches the next command from the user.
func (s *CommandReader) ReadCommand() rslt.Of[core.Command] {
	fmt.Print(">>> ")
	line, err := s.reader.ReadString('\n')
	if err != nil {
		return rslt.Error[core.Command](err)
	}

	return rslt.Value(core.Command(strings.TrimSpace(line)))
}
