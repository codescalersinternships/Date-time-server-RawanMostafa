package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)
const portNum string = ":8080"

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my datetime server!")
}

func getDate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		currentTime := time.Now()
		fmt.Fprint(w, currentTime.String())
	}
}

func main() {
	fmt.Println("Starting our simple http server.")

	http.HandleFunc("/", home)
	http.HandleFunc("/datetime", getDate)

	fmt.Println("Started on port", portNum)
	fmt.Println("To close connection CTRL+C :-)")

	err := http.ListenAndServe(portNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
