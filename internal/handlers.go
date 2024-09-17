package internal 

import (
	"fmt"
	"net/http"
	"time"
)

const PortNum string = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my datetime server!")
}

func GetDate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		currentTime := time.Now()
		fmt.Fprint(w, currentTime.String())
	}
}