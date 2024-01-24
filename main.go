package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	go func() {
		for {
			item := <-channel
			fmt.Println(item)
		}
	}()
}

var channel = make(chan int)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		channel <- 3
		ctx.String(200, "Hello Gin")
	})

	router.Run("localhost:8080")
}
