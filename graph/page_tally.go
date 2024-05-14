package graph

import (
	"os"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/log"
)

type PageTally struct {
	PublicPageCount  int
	PrivatePageCount int
	TotalPageCount   int
}

// NewPageTally creates a new PageTally struct.
func NewPageTally(pages []*Page) *PageTally {
	publicPageCount := 0
	for _, page := range pages {
		if page.IsPublic() {
			publicPageCount++
		}
	}
	totalPageCount := len(pages)
	privatePageCount := totalPageCount - publicPageCount
	log.Debug("Tallied pages", "public", publicPageCount, "private", privatePageCount, "total", totalPageCount)

	return &PageTally{
		PublicPageCount:  publicPageCount,
		PrivatePageCount: privatePageCount,
		TotalPageCount:   totalPageCount,
	}
}

// Render renders a table tallying the number of public and private pages.
func (pt PageTally) Render() string {
	const (
		purple    = lipgloss.Color("99")
		gray      = lipgloss.Color("245")
		lightGray = lipgloss.Color("251")
	)
	re := lipgloss.NewRenderer(os.Stdout)
	var (
		// HeaderStyle is the lipgloss style used for the table headers.
		HeaderStyle = re.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)
		// CellStyle is the base lipgloss style used for the table rows.
		CellStyle = re.NewStyle().Padding(0, 1).Width(14)
		// RowStyle is the lipgloss style used for regular table rows.
		RowStyle = CellStyle.Foreground(lightGray)
		// BorderStyle is the lipgloss style used for the table border.
		BorderStyle = lipgloss.NewStyle().Foreground(purple)
	)
	rows := [][]string{
		{"Private", strconv.Itoa(pt.PrivatePageCount)},
		{"Public", strconv.Itoa(pt.PublicPageCount)},
		{"Total", strconv.Itoa(pt.TotalPageCount)},
	}
	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(BorderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return HeaderStyle
			default:
				return RowStyle
			}
		}).
		Headers("Pages", "Count").
		Rows(rows...)

	return t.Render()
}
