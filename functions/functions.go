package functions

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func MdToHTML(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func EntryDate(dirPath string) (string, error) {
        var entries []string
    err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
            createdAt := info.ModTime().Format("Jan 2, 2006")
            entry := fmt.Sprintf("%s (%s)", strings.TrimSuffix(info.Name(), ".md"), createdAt)
            entries = append(entries, entry)
        }
        return nil
    })
    if err != nil {
        return "", err
    }
    return strings.Join(entries, "\n"), nil
}

func ListMarkdownFiles(dirPath string) ([]string, error) {
    var entries []string

    err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
            entries = append(entries, strings.TrimSuffix(info.Name(), ".md"))
        }

        return nil
    })

    if err != nil {
        return nil, err
    }

    return entries, nil
}


// Format filename. 
func DocumentFormatter(filepath string) string {
    var outputFormat string = "YYYY-DD-MM_[ENTRY_TITLE].md"
    return outputFormat
}

// Send MD document to filesystem.
func WriteEntry (){
}
