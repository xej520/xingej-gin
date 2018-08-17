package querystringandbody

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test_query_a(t *testing.T)  {
	router := gin.Default()

	router.GET("/welcome", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "Guest")
		lastname := context.Query("lastname")

		// 使用String类型，作为返回类型
		context.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Run()

}


