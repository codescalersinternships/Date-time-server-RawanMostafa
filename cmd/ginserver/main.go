package main

import (
	"log"

	"github.com/codescalersinternships/Datetime-server-RawanMostafa/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", pkg.GinHome)
	r.GET("/datetime", pkg.GinHandler)
	err := r.Run(":8083")
	if err != nil {
		log.Fatalf("Error: impossible to start server: %s", err)
	}
}
