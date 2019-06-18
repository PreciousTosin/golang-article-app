package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
	// Respond with HTML
	c.HTML(http.StatusOK, templateName, data)
}
