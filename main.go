package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "voila",
		})
	})
	return r
}

func main() {
	r := setupRouter()

	r.Run() // listen and serve on 8080
}
