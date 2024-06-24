package model

import (
	"html/template"
	"net/http"

	"github.com/abe444/BUBBLE_ENGINE/functions"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			c.HTML(http.StatusOK, "userPanel.html", gin.H{
				"headerTags": template.HTML(functions.DisplayHead()),
				"title":      "USER_PANEL_TITLE_CONFIG",
				"message":    "You are not authorized to view this page. Please log in.",
			})
			return
		}
		c.Next()
	}
}
