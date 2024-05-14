package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"

	"github.com/brianwisti/rgb_logseq/page"
	"github.com/brianwisti/rgb_logseq/reports"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	graphDir := os.Getenv("GRAPH_DIR")
	pagesDir := filepath.Join(graphDir, "pages")
	log.Debug(pagesDir)
	sourcePaths, err := filepath.Glob(filepath.Join(pagesDir, "*.md"))
	if err != nil {
		log.Fatal(err)
	}

	pages := []*page.Page{}

	for _, sourceFile := range sourcePaths {
		p := page.NewPage(sourceFile)
		pages = append(pages, p)
	}

	pt := reports.NewPageTally(pages)
	fmt.Println(pt.Render())
}
