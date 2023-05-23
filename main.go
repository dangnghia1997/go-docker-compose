package main

import (
	"blog/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	infrastructure.LoadEnv()     //loading env
	infrastructure.NewDatabase() //new database connection
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "Hello World 222 !"})
	})
	router.Run(":8000")
}
