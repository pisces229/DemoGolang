package sample

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Redirect(ginEngine *gin.Engine) {
	ginEngine.GET("/redirect1", func(ginContext *gin.Context) {
		ginContext.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})
	ginEngine.GET("/redirect2", func(ginContext *gin.Context) {
		ginContext.Redirect(http.StatusFound, "/redirect1")
	})
}
