package ginhandler

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GinHandler(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.String(http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	formattedTime := time.Now().Format(time.ANSIC)

	if strings.Contains(c.Request.Header.Get("content-type"), "text/plain") {
		c.Writer.Header().Set("Content-Type", "text/plain")
		c.String(http.StatusOK, formattedTime)

	} else if strings.Contains(c.Request.Header.Get("content-type"), "application/json") {

		c.JSON(http.StatusOK, formattedTime)
	} else {
		c.String(http.StatusUnsupportedMediaType, http.StatusText(http.StatusUnsupportedMediaType))
	}

}

func GinHome(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to my datetime server!")
}
