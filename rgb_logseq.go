package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"

	"github.com/brianwisti/rgb_logseq/graph"
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

	pages := []*graph.Page{}

	for _, sourceFile := range sourcePaths {
		p := graph.NewPage(sourceFile)
		pages = append(pages, p)
	}

	pt := graph.NewPageTally(pages)
	fmt.Println(pt.Render())
}
