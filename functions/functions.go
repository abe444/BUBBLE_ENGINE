package functions

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/abe444/BUBBLE_ENGINE/types"
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

func ListMarkdownFiles(dirPath string) ([]string, error) {
	var entries []string
	fileInfos := make([]types.FileInfo, 0)

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			fileInfos = append(fileInfos, types.FileInfo{
				Name:    strings.TrimSuffix(info.Name(), ".md"),
				ModTime: info.ModTime(),
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].ModTime.After(fileInfos[j].ModTime)
	})

	for _, fi := range fileInfos {
		entries = append(entries, fi.Name)
	}

	return entries, nil
}
