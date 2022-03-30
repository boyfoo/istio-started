package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	engine := gin.New()
	engine.GET("/pords/:id", func(ctx *gin.Context) {
		query := ctx.Param("id")
		getenv := os.Getenv("REVIEW_IP")
		req := "127.0.0.1:8081"
		if getenv != "" {
			req = getenv
		}
		res, err := http.Get("http://" + req + "/review/" + query)
		if err != nil {
			ctx.JSON(417, gin.H{"err": err})
			return
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			ctx.JSON(417, gin.H{"err": err})
			return
		}
		ctx.JSON(200, gin.H{
			"id": string(body),
		})
	})
	engine.Run(":8080")
}
