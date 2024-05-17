package graph

import "testing"

func TestLinesInNewBlock(t *testing.T) {
	tests := []struct {
		lines []string
		want  int
	}{
		{[]string{""}, 1},
		{[]string{"Hello World"}, 1},
		{[]string{"title:: Test Page"}, 0},
		{[]string{"title:: Test Page", "Hello World"}, 1},
	}

	for _, tt := range tests {
		logseqLines := make([]*LogseqLine, 0)

		for _, line := range tt.lines {
			l, _ := NewLogseqLine(line)
			logseqLines = append(logseqLines, l)
		}

		b, err := NewBlock(logseqLines)

		if err != nil {
			t.Errorf("NewBlock(%q) error = %v; want nil", tt.lines, err)
		}

		got := len(b.Lines)
		if got != tt.want {
			t.Errorf("NewBlock(%q) = %d; want %d", tt.lines, got, tt.want)
		}
	}
}

func TestBlockDepth(t *testing.T) {
	tests := []struct {
		lines    []string
		expected int
	}{
		{[]string{"Hello World"}, 0},
		{[]string{"- Hello World"}, 1},
		{[]string{"  Hello World"}, 1},
		{[]string{"\t- Hello World"}, 2},
		{[]string{"\t  Hello World"}, 2},
		{[]string{"\t  Hello World", "\t  Hello World"}, 2},
		{[]string{"\t\t-  Hello World", "\t\t  Hello World"}, 3},
	}

	for _, tt := range tests {
		logseqLines := make([]*LogseqLine, 0)

		for _, line := range tt.lines {
			l, _ := NewLogseqLine(line)
			logseqLines = append(logseqLines, l)
		}

		b, _ := NewBlock(logseqLines)
		got := b.Depth

		if got != tt.expected {
			t.Errorf("NewBlock(%q) = %d; want %d", tt.lines, got, tt.expected)
		}
	}
}

func TestErrMismatchedDepth(t *testing.T) {
	tests := []struct {
		lines []string
	}{
		{[]string{"Hello World", "  Hello World"}},
		{[]string{"- Hello World", "\t  Hello World"}},
		{[]string{"\t  Hello World", "- Hello World"}},
	}

	for _, tt := range tests {
		logseqLines := make([]*LogseqLine, 0)

		for _, line := range tt.lines {
			l, _ := NewLogseqLine(line)
			logseqLines = append(logseqLines, l)
		}

		_, err := NewBlock(logseqLines)

		if err == nil {
			t.Errorf("NewBlock(%q) error = nil; want %v", tt.lines, ErrMismatchedDepth)
		}
	}
}

func TestGetProp(t *testing.T) {
	tests := []struct {
		lines []string
		field string
		want  string
	}{
		{[]string{"title:: Test Page"}, "title", "Test Page"},
		{[]string{"title:: Test Page", "public:: true"}, "public", "true"},
		{[]string{"title:: Test Page", "Hello World"}, "title", "Test Page"},
	}

	for _, tt := range tests {
		logseqLines := make([]*LogseqLine, 0)

		for _, line := range tt.lines {
			l, _ := NewLogseqLine(line)
			logseqLines = append(logseqLines, l)
		}

		b, _ := NewBlock(logseqLines)
		got, err := b.GetProp(tt.field)

		if err != nil {
			t.Errorf("GetProp(%q) error = %v; want nil", tt.lines, err)
		}

		if got.AsString() != tt.want {
			t.Errorf("GetProp(%q) = %q; want %q", tt.lines, got, tt.want)
		}
	}
}
