package main

import (
	"fmt"
	"log"
	"net/http"
	internal "github.com/codescalersinternships/Datetime-server-RawanMostafa/internal"
)

func main() {
	fmt.Println("Starting our simple http server.")

	http.HandleFunc("/", internal.HttpHome)
	http.HandleFunc("/datetime", internal.HttpHandler)

	fmt.Println("Started on port", internal.PortNum)

	err := http.ListenAndServe(internal.PortNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
