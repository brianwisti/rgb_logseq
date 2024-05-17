package graph

import (
	"errors"
	"strings"
)

const (
	// BlockPrefix is the prefix for a block opener.
	BlockPrefix = "- "
	// BlockContentPrefix is the prefix for a continued block content line.
	BlockContentPrefix = "  "
	// IndentMarker is the character used to indicate indentation.
	IndentMarker = "\t"
	// PropIndicator helps identify properties.
	PropIndicator = ":: "
)

// ErrInvalidLogseqLine is returned when a line is not a valid Logseq line.
var ErrInvalidLogseqLine = errors.New("invalid Logseq line")

// LogseqLine holds the data for a single line in a Logseq page.
type LogseqLine struct {
	// Raw is the raw text of the line.
	Raw string
	// Adjusted is the line with outline structures removed.
	Adjusted string
	// Depth is the indentation level of the line.
	Depth int
	// IsProp is true if the line indicates a property.
	IsProp bool
}

// NewLogseqLine creates a new LogseqLine.
func NewLogseqLine(line string) (*LogseqLine, error) {
	l := &LogseqLine{Raw: line}
	err := l.adjust()
	return l, err
}

// adjust sets the adjusted line and depth.
func (l *LogseqLine) adjust() error {
	l.Adjusted = l.Raw
	l.Depth = 0

	for strings.HasPrefix(l.Adjusted, IndentMarker) {
		l.Adjusted = strings.TrimPrefix(l.Adjusted, IndentMarker)
		l.Depth++
	}

	// Return an error if the line has depth but no block markers
	hasBlockMarker := strings.HasPrefix(l.Adjusted, BlockPrefix) || strings.HasPrefix(l.Adjusted, BlockContentPrefix)

	if l.Depth > 0 && !hasBlockMarker {
		return ErrInvalidLogseqLine
	}

	l.Adjusted = strings.TrimPrefix(l.Adjusted, BlockPrefix)
	l.Adjusted = strings.TrimPrefix(l.Adjusted, BlockContentPrefix)
	l.IsProp = strings.Contains(l.Adjusted, PropIndicator)

	return nil
}
