package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*func getCookie(req *http.Request) string {
	cookie, _ := req.Cookie("token")
	fmt.Println("Cookie", cookie)
	// return cookie
}*/

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
	// var req http.Request
	loggedInInterface, _ := c.Get("is_logged_in")
	// cookie, _ := req.Cookie("token")
	if loggedInInterface != nil {
		data["is_logged_in"] = loggedInInterface.(bool)
	}
	// data["is_logged_in"] = loggedInInterface.(bool)
	fmt.Println("Interface", loggedInInterface)
	// fmt.Println("Cookie", cookie)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
