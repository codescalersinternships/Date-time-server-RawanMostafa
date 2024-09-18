package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
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
		return
	}
	currentTime := time.Now()
	trucatedTime := truncateToSec(currentTime)

	if strings.Contains(r.Header.Get("content-type"), "text/plain") {

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, trucatedTime.String())

	} else if strings.Contains(r.Header.Get("content-type"), "json") {

		w.Header().Set("Content-Type", "application/json")

		timeJson, err := json.Marshal(trucatedTime)
		if err != nil {
			log.Fatalf("error converting to json: %v", err)
		}
		_, err = w.Write(timeJson)
		if err != nil {
			log.Fatalf("error writing data to response: %v", err)
		}
	}

}

func GinHandler(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.String(http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	currentTime := time.Now()
	trucatedTime := truncateToSec(currentTime)

	if strings.Contains(c.Request.Header.Get("content-type"), "text/plain") {
		c.Writer.Header().Set("Content-Type", "text/plain")
		c.String(http.StatusOK, trucatedTime.String())

	} else if strings.Contains(c.Request.Header.Get("content-type"), "json") {

		c.JSON(http.StatusOK, trucatedTime)
	}

}

func GinHome(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to my datetime server!")
}
