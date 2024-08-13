package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name": "Egi",
			"desc": "intern",
			"time": "10 month",
		})
	})

	r.GET("/path2", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"divisi":   "Engineer",
			"position": "Backend",
			"materi":   "Belajar API pada golang",
		})
	})

	r.Run(":8888")
}
