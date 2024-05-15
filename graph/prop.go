package graph

import (
	"errors"
	"strings"
)

var ErrInvalidLine = errors.New("invalid line")
var ErrPropNotBoolean = errors.New("property is not a boolean")

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

// AsString returns the original string value of the property.
func (p *Prop) AsString() string {
	return p.value
}

// AsBoolean returns the value of the property as a boolean, or false with an ErrPropNotBoolean if the value is not a boolean.
func (p *Prop) AsBoolean() (bool, error) {
	if p.value == "true" {
		return true, nil
	}

	if p.value == "false" {
		return false, nil
	}

	return false, ErrPropNotBoolean
}
