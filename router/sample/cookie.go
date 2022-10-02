package sample

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Cookie(ginEngine *gin.Engine) {
	ginEngine.GET("/cookie", func(ginContext *gin.Context) {
		cookie, err := ginContext.Cookie("cookie")
		if err != nil {
			cookie = "NotSet"
			ginContext.SetCookie("cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	})
}
