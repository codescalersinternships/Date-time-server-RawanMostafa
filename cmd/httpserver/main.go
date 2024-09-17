package main

import (
	"fmt"
	"log"
	"net/http"
	internal "github.com/codescalersinternships/Datetime-server-RawanMostafa/internal"
)

func main() {
	fmt.Println("Starting our simple http server.")

	http.HandleFunc("/", internal.Home)
	http.HandleFunc("/datetime", internal.GetDate)

	fmt.Println("Started on port", internal.PortNum)
	fmt.Println("To close connection CTRL+C :-)")

	err := http.ListenAndServe(internal.PortNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
