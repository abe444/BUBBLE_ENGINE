package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/abe444/BUBBLE_ENGINE/functions"
	"github.com/abe444/BUBBLE_ENGINE/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func StartupRoutes(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {

		entries, err := functions.ListMarkdownFiles("./entries")
		if err != nil {
			log.Fatal(err)
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"headerTags": template.HTML(functions.DisplayHead()),
			"title":      "TITLE_CONFIG",
			"entries":    entries,
		})
	})

	router.GET("/login", gin.BasicAuth(gin.Accounts{
		"abe": "pass123",
	}), func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		session := sessions.Default(c)
		session.Set("user", user)
		session.Save()
		c.HTML(http.StatusOK, "login.html", gin.H{
			"headerTags": template.HTML(functions.DisplayHead()),
			"title":      "TITLE_CONFIG",
			"user":       user,
		})
	})

	authorized := router.Group("/")
	authorized.Use(model.AuthRequired())
	{
		authorized.GET("/panel", func(c *gin.Context) {
			user := sessions.Default(c).Get("user")
			c.HTML(http.StatusOK, "userPanel.html", gin.H{
				"headerTags": template.HTML(functions.DisplayHead()),
				"title":      "USER_PANEL_TITLE_CONFIG",
				"user":       user,
			})
		})
	}

	router.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		seconds := 3
		time.Sleep(time.Duration(seconds) * time.Second)
		c.Redirect(http.StatusFound, "/")
	})

	router.GET("/blog", func(c *gin.Context) {
		entries, err := functions.ListMarkdownFiles("./entries")
		if err != nil {
			log.Fatal(err)
		}

		c.HTML(http.StatusOK, "blog.html", gin.H{
			"headerTags": template.HTML(functions.DisplayHead()),
			"title":      "TITLE_CONFIG",
			"entries":    entries,
		})
	})

	router.POST("/submit_article", func(c *gin.Context) {
		file, err := c.FormFile("submit_article")
		if err != nil {
			c.String(http.StatusBadRequest, "error %v", err)
			return
		}

		filePath := filepath.Join("./entries", file.Filename)
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error saving file: %v", err)
			return
		}
		c.Redirect(http.StatusFound, "/")
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
		html := functions.MdToHTML(md)

		fileInfo, err := os.Stat(path)
		if err != nil {
			fmt.Println("Error getting file info:", err)
			return
		}

		modifiedTime := fileInfo.ModTime()
		formattedModifiedTimes := make([]string, 1)
		formattedModifiedTimes[0] = modifiedTime.Format("January 2, 2006")

		c.HTML(http.StatusOK, "entry.html", gin.H{
			"headerTags": template.HTML(functions.DisplayHead()),
			"title":      "TITLE_CONFIG",
			"created":    formattedModifiedTimes,
			"modified":   formattedModifiedTimes,
			"entries":    template.HTML(html),
			"donate":     "xmr addy here",
		})
	})

	router.Run(":8080")
}
