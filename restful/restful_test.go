package restful

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)
// 测试冒号
//冒号:加上一个参数名组成路由参数
func Test_restful_A(t *testing.T)  {
	router := gin.Default()

	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "Hello %s", name)
	})

	router.Run(":8080")
}

func Test_restful_B(t *testing.T)  {
	router := gin.Default()

	router.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")

		message := name +" is " + action

		context.String(http.StatusOK, message)
	})

	router.Run(":8080")
}


