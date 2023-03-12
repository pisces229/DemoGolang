package sample

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler(ginEngine *gin.Engine) {
	ginEngine.GET("/handler", func(ginContext *gin.Context) {
		ginContext.String(http.StatusOK, "[1]")
		if true {
			ginContext.Abort()
		} else {
			//ginContext.Next()
		}
		ginContext.String(http.StatusOK, "[2]")
	}, func(ginContext *gin.Context) {
		ginContext.String(http.StatusOK, "[3]")
	})
}
