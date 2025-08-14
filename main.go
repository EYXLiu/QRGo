package main

import (
	"github.com/gin-gonic/gin"
	"QRGo/handlers"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handlers.PingHandler)
	r.GET("/qr", handlers.QRHandler)
    r.Run(":8080")
}
