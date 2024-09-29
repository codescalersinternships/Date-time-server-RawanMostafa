package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codescalersinternships/Datetime-server-RawanMostafa/pkg"
)

func main() {
	fmt.Println("Starting our simple http server.")

	http.HandleFunc("/", pkg.HttpHome)
	http.HandleFunc("/datetime", pkg.HttpHandler)

	fmt.Println("Started on port", ":8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
