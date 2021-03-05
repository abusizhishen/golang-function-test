package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.POST("/b", func(context *gin.Context) {
		fmt.Println(context.GetHeader("name"))
		fmt.Println(context.GetHeader("name"))
	})

	g.Run(":90")
}