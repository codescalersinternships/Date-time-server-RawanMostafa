package internal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const PortNum string = ":8080"

func truncateToSec(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
}

func HttpHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my datetime server!")
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	} else {
		currentTime := time.Now()
		trucatedTime := truncateToSec(currentTime)
		fmt.Fprint(w, trucatedTime.String())
	}
}

func GinHandler(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.String(http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	currentTime := time.Now()
	trucatedTime := truncateToSec(currentTime)
	c.String(http.StatusOK, trucatedTime.String())

}

func GinHome(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to my datetime server!")
}
