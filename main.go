package main

import (
	"net/http"
	"os"
	"path/filepath"
    "html/template"

	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func mdToHTML(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func main() {
    router := gin.Default()

    router.LoadHTMLGlob("templates/*")

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    router.GET("/blog/:post", func(c *gin.Context) {
        post := c.Param("post")
        path := filepath.Join("entries", post+".md")
        entries, err := os.ReadFile(path)
        if err != nil {
            c.String(http.StatusNotFound, "Post not found")
            return
        }

        md := entries
        html := mdToHTML(md)
        c.HTML(http.StatusOK, "entry.html", gin.H{
            "Entries": template.HTML(html),
        })
    })

    router.Run(":8080")
}
