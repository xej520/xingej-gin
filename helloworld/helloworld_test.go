package helloworld

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func Test_helloworld(t *testing.T)  {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello world, gin")
	})

	router.Run(":8080")
}


