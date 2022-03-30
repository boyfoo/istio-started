package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.GET("/review/:id", func(ctx *gin.Context) {
		query := ctx.Param("id")
		fmt.Println(query)
		ctx.JSON(200, gin.H{
			"review": query + " 今天天气不错",
		})
	})
	engine.Run(":8081")
}
