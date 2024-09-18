package main

import (
	"fmt"
	"log"
	"net/http"

	pkg "github.com/codescalersinternships/Datetime-server-RawanMostafa/pkg"
)

func main() {
	fmt.Println("Starting our simple http server.")

	http.HandleFunc("/", pkg.HttpHome)
	http.HandleFunc("/datetime", pkg.HttpHandler)

	fmt.Println("Started on port", pkg.PortNum)

	err := http.ListenAndServe(pkg.PortNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
