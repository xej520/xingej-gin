package querystringandbody

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http"
)
//常见的格式就有四种。
// application/json，
// application/x-www-form-urlencoded,
// application/xml
// multipart/form-data
//后面一个主要用于图片上传。
// json格式的很好理解，
// urlencode其实也不难，无非就是把query string的内容，放到了body体里，同样也需要urlencode。
// 默认情况下，c.PostFROM解析的是x-www-form-urlencoded或from-data的参数。


func Test_body_a(t *testing.T)  {
	router := gin.Default()

	router.POST("/form_post", func(context *gin.Context) {
		message := context.PostForm("message")
		nick := context.DefaultPostForm("nick", "anonymous")

		context.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"status_code": http.StatusOK,
				"status": "ok",
			},
			"message" : message,
			"nick": nick,
		})
	})

	// 默认是8080端口
	router.Run()

}

