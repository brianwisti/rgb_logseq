package graph

import (
	"errors"

	"github.com/charmbracelet/log"
)

var ErrEmptyBlock = errors.New("empty block")
var ErrMismatchedDepth = errors.New("mismatched depth")
var ErrPropNotFound = errors.New("property not found")

// Block represents a single block in a Logseq page.
type Block struct {
	// Lines is the content of the block.
	Lines []*LogseqLine
	// Props is a map of properties for the block.
	Props map[string]*Prop
	// Depth is the indentation level of the block.
	Depth int
}

// NewBlock creates a new block from a slice of lines.
func NewBlock(lines []*LogseqLine) (*Block, error) {

	if len(lines) == 0 {
		return nil, ErrEmptyBlock
	}

	b := &Block{Lines: lines}
	err := b.setDepth()

	if err != nil {
		return nil, err
	}

	b.setProps()
	b.stripProps()
	return b, nil
}

// setDepth sets the depth of the block.
func (b *Block) setDepth() error {
	b.Depth = b.Lines[0].Depth

	for _, l := range b.Lines {
		if l.Depth != b.Depth {
			log.Error("Mismatched depth", "line", l.Adjusted, "depth", l.Depth, "block", b.Depth)
			return ErrMismatchedDepth
		}
	}

	return nil
}

// setProps sets the properties of the block.
func (b *Block) setProps() {
	b.Props = make(map[string]*Prop)

	for _, l := range b.Lines {
		if l.IsProp {
			prop, err := NewProp(l.Adjusted)
			if err == nil {
				b.Props[prop.field] = prop
			}
		}
	}
}

// stripProps removes property lines from the block.
func (b *Block) stripProps() {
	var newLines []*LogseqLine

	for _, l := range b.Lines {
		if !l.IsProp {
			newLines = append(newLines, l)
		}
	}

	b.Lines = newLines
}

// GetProp returns a property from the block.
func (b *Block) GetProp(field string) (*Prop, error) {
	prop, ok := b.Props[field]

	if !ok {
		return nil, ErrPropNotFound
	}

	return prop, nil
}
