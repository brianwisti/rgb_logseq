package graph

import "testing"

func TestLineTrimming(t *testing.T) {
	tests := []struct {
		line     string
		expected string
	}{
		{"", ""},
		{"Hello World", "Hello World"},
		{"- Hello World", "Hello World"},
		{"  Hello World", "Hello World"},
		{"\t- Hello World", "Hello World"},
		{"\t  Hello World", "Hello World"},
		{"title:: Test Page", "title:: Test Page"},
		{"  title:: Test Page", "title:: Test Page"},
		{"\t  title:: Test Page", "title:: Test Page"},
	}

	for _, tt := range tests {
		l, err := NewLogseqLine(tt.line)

		if err != nil {
			t.Errorf("NewLogseqLine(%q) error = %v; want nil", tt.line, err)
		}

		if l.Adjusted != tt.expected {
			t.Errorf("NewLogseqLine(%q) = %q; want %q", tt.line, l.Adjusted, tt.expected)
		}
	}
}

func TestErrInvalidLogseqLine(t *testing.T) {
	tests := []string{
		"\ttitle:: Test Page",
		"\tinvalid line",
	}

	for _, tt := range tests {
		_, err := NewLogseqLine(tt)

		if err == nil {
			t.Errorf("NewLogseqLine(%q) error = nil; want %v", tt, ErrInvalidLogseqLine)
		}
	}

}

func TestDepth(t *testing.T) {
	tests := []struct {
		line     string
		expected int
	}{
		{"", 0},
		{"Hello World", 0},
		{"- Hello World", 1},
		{"  Hello World", 1},
		{"\t- Hello World", 2},
		{"\t  Hello World", 2},
		{"\t\t- Hello World", 3},
		{"\t\t  Hello World", 3},
	}

	for _, tt := range tests {
		l, _ := NewLogseqLine(tt.line)

		if l.Depth != tt.expected {
			t.Errorf("NewLogseqLine(%q) depth = %d; want %d", tt.line, l.Depth, tt.expected)
		}
	}
}

func TestIsProp(t *testing.T) {
	tests := []struct {
		line     string
		expected bool
	}{
		{"title:: Test Page", true},
		{"  title:: Test Page", true},
		{"\t  title:: Test Page", true},
		{"title :: Test Page", true}, // not 100% sure about this one
		{"title: Test Page", false},
		{"title is Test Page", false},
	}

	for _, tt := range tests {
		l, _ := NewLogseqLine(tt.line)

		if l.IsProp != tt.expected {
			t.Errorf("NewLogseqLine(%q) IsProp = %t; want %t", tt.line, l.IsProp, tt.expected)
		}
	}
}
