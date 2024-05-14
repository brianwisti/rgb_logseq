package graph

import "testing"

func TestLineIsProp(t *testing.T) {
	tests := []struct {
		line string
		want bool
	}{
		{"title:: Test Page", true},
		{"", false},
		{"- ", false},
		{"# ", false},
		{"- title:: Test Page", false},
	}

	for _, tt := range tests {
		got := LineIsProp(tt.line)
		if got != tt.want {
			t.Errorf("LineIsProp(%q) = %t; want %t", tt.line, got, tt.want)
		}
	}
}

func TestNewProp(t *testing.T) {
	tests := []struct {
		input       string
		wantedField string
		wantedValue string
	}{
		{"title:: Test Page", "title", "Test Page"},
		{"public:: true", "public", "true"},
	}

	for _, tt := range tests {
		p, err := NewProp(tt.input)

		if err != nil {
			t.Errorf("NewProp(%q) error = %v; want nil", tt.input, err)
		}

		if p.field != tt.wantedField {
			t.Errorf("NewProp(%q) field = %q; want %q", tt.input, p.field, tt.wantedField)
		}
		if p.value != tt.wantedValue {
			t.Errorf("NewProp(%q) value = %q; want %q", tt.input, p.value, tt.wantedValue)
		}
	}
}

func TestInvalidLine(t *testing.T) {
	tests := []string{
		"- title:: Test Page",
		"invalid line",
	}

	for _, tt := range tests {
		p, err := NewProp(tt)

		if err != ErrInvalidLine {
			t.Errorf("NewProp(%q) error = %v; want ErrInvalidLine", tt, err)
		}

		if p != nil {
			t.Errorf("NewProp(%q) = %v; want nil", tt, p)
		}
	}
}
