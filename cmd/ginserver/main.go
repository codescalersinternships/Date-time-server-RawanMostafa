package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	ginhandler "github.com/codescalersinternships/Datetime-server-RawanMostafa/pkg/ginserver"

	"github.com/gin-gonic/gin"
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
	r := gin.Default()
	r.GET("/", ginhandler.GinHome)
	r.GET("/datetime", ginhandler.GinHandler)
	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Error: impossible to start server: %s", err)
	}
}
