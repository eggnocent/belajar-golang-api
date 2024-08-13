package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/rootHandler", rootHandler)

	r.GET("/helloHandler", helloHandler)

	r.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Egi",
		"desc": "intern",
		"time": "10 month",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"divisi":   "Engineer",
		"position": "Backend",
		"materi":   "Belajar API pada golang",
	})
}
