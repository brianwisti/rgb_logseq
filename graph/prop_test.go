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

func TestAsString(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"title:: Test Page", "Test Page"},
		{"public:: true", "true"},
	}

	for _, tt := range tests {
		p, _ := NewProp(tt.input)
		got := p.AsString()

		if got != tt.want {
			t.Errorf("Prop.AsString() = %q; want %q", got, tt.want)
		}
	}
}

func TestAsBoolean(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		p := Prop{value: tt.input}
		got, err := p.AsBoolean()

		if err != nil {
			t.Errorf("Prop.AsBoolean() error = %v; want nil", err)
		}

		if got != tt.want {
			t.Errorf("Prop.AsBoolean() = %t; want %t", got, tt.want)
		}
	}
}

func TestErrPropNotBoolean(t *testing.T) {
	tests := []string{
		"invalid",
		"true ",
		" false",
	}

	for _, tt := range tests {
		p := Prop{value: tt}
		got, err := p.AsBoolean()
		if got {
			t.Errorf("Prop.AsBoolean() = %t; want false", got)
		}

		if err != ErrPropNotBoolean {
			t.Errorf("Prop.AsBoolean() error = %v; want ErrPropNotBoolean", err)
		}
	}
}
