package internal

import (
	"fmt"
	"net/http"
	"time"
)

const PortNum string = ":8080"

func truncateToSec(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my datetime server!")
}

func GetDate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		currentTime := time.Now()
		trucatedTime := truncateToSec(currentTime)
		fmt.Fprint(w, trucatedTime.String())
	}
}
