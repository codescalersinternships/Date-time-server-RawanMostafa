package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	httphandler "github.com/codescalersinternships/Datetime-server-RawanMostafa/pkg/httpserver"
)

const defaultPort = "8083"

func getFlags() string {
	var port string
	flag.StringVar(&port, "port", "", "Specifies the port")

	flag.Parse()
	return port
}

func decideConfigs() string {

	port := getFlags()

	if port == "" {
		envPort, found := os.LookupEnv("DATETIME_PORT")

		if found {
			port = envPort
		} else {
			port = defaultPort
		}
	}
	return port

}
func main() {
	port := decideConfigs()

	fmt.Println("Starting our simple http server.")

	http.HandleFunc("/", httphandler.HttpHome)
	http.HandleFunc("/datetime", httphandler.HttpHandler)

	fmt.Printf("Started on port :%s\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
