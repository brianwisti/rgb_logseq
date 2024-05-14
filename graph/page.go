package graph

import (
	"bufio"
	"os"
	"strings"

	"github.com/charmbracelet/log"
)

type Page struct {
	sourceFile string
	props      map[string]string
}

func NewPage(sourceFile string) *Page {
	log.Debug("Reading page", "source", sourceFile)
	readFile, err := os.Open(sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	defer readFile.Close()
	scanner := bufio.NewScanner(readFile)
	pageProps := make(map[string]string)
	inProps := true

	for scanner.Scan() {
		line := scanner.Text()

		// Stop reading when we reach an empty line.
		if line == "" {
			inProps = false
		}

		if strings.HasPrefix(line, "-") {
			inProps = false
		}

		if strings.HasPrefix(line, "#") {
			inProps = false
		}

		if inProps {
			if strings.Index(line, "::") > 0 {
				// split the line into field and value from ":: "
				fields := strings.SplitN(line, ":: ", 2)
				if len(fields) != 2 {
					log.Fatalf("invalid line %q", line)
				}

				field, value := fields[0], fields[1]
				pageProps[field] = value
			}
		}
	}
	return &Page{sourceFile: sourceFile, props: pageProps}
}

func (p *Page) IsPublic() bool {
	return p.props["public"] == "true"
}
