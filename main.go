package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	go func() {
		for {
			now := <-channel
			fmt.Println(now)
		}
	}()
}

var channel = make(chan time.Time)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		channel <- time.Now()
		ctx.String(200, "Hello Gin")
	})

	router.Run("localhost:8080")
}
