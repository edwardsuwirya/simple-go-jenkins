package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	log.Println("Hello World")
	run(port)
}

func initRouter(route *gin.Engine) {
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func run(port string) {
	route := gin.Default()
	initRouter(route)
	err := route.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Server is not running %s", err)
	}
}
