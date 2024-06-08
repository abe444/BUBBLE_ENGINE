package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

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
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./favicon.ico")


	var entries []string

	err := filepath.Walk("entries", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			entries = append(entries, strings.TrimSuffix(info.Name(), ".md"))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "TITLE_CONFIG",
			"entries": entries,
		})
	})

	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", gin.H{
			"title":   "TITLE_CONFIG",
		})
	})

	router.GET("/blog", func(c *gin.Context) {
        created := time.Now()
		c.HTML(http.StatusOK, "blog.html", gin.H{
			"title":   "TITLE_CONFIG",
            "created": created.Format("Jan 2, 2006"), 
			"entries": entries,
		})
	})

	router.GET("/entry/:post", func(c *gin.Context) {
		post := c.Param("post")
		path := filepath.Join("entries", post+".md")
		entries, err := os.ReadFile(path)
		if err != nil {
			c.String(http.StatusNotFound, "Post not found")
			return
		}

		md := entries
		html := mdToHTML(md)

        fileInfo, err := os.Stat(path)
        if err != nil {
            fmt.Println("Error getting file info:", err)
            return
        }

        modifiedTime := fileInfo.ModTime()
        formattedModifiedTimes := make([]string, 1)
        formattedModifiedTimes[0] = modifiedTime.Format("January 2, 2006") 

		c.HTML(http.StatusOK, "entry.html", gin.H{
			"title":   "TITLE_CONFIG",
            "created": formattedModifiedTimes,
            "modified": formattedModifiedTimes,
			"entries":  template.HTML(html),
            "donate": "xmr addy here",
		})
	})

	router.Run(":8080")
}
