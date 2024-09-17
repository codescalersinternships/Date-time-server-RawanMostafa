package main

import (
	"log"

	internal "github.com/codescalersinternships/Datetime-server-RawanMostafa/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/",internal.GinHome)
	r.GET("/datetime", internal.GinHandler)
	err := r.Run(":8083")
	if err != nil {
		log.Fatalf("Error: impossible to start server: %s", err)
	}
}
