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

func HttpHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my datetime server!")
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	currentTime := time.Now()
	formattedTime := currentTime.Format(time.ANSIC)

	if strings.Contains(r.Header.Get("content-type"), "text/plain") {

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, formattedTime)

	} else if strings.Contains(r.Header.Get("content-type"), "application/json") {

		w.Header().Set("Content-Type", "application/json")

		timeJson, err := json.Marshal(formattedTime)
		if err != nil {
			log.Fatalf("error converting to json: %v", err)
		}
		_, err = w.Write(timeJson)
		if err != nil {
			log.Fatalf("error writing data to response: %v", err)
		}
	} else {
		http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
	}

}

func GinHandler(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.String(http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	currentTime := time.Now()
	formattedTime := currentTime.Format(time.ANSIC)

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
