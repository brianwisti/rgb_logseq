package graph

import (
	"errors"
	"strings"
)

var ErrInvalidLine = errors.New("invalid line")

// Prop represents a Logseq property.
type Prop struct {
	field string
	value string
}

func NewProp(line string) (*Prop, error) {

	if !LineIsProp(line) {
		return nil, ErrInvalidLine
	}

	fields := strings.SplitN(line, ":: ", 2)
	if len(fields) != 2 {
		return nil, ErrInvalidLine
	}

	return &Prop{field: fields[0], value: fields[1]}, nil
}

// LineIsProp returns true if the line describes a Logseq property.
func LineIsProp(line string) bool {
	if strings.Index(line, "-") == 0 {
		return false
	}

	return strings.Index(line, "::") > 0
}
