package sample

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Middleware(ginEngine *gin.Engine) {
	defaultMiddleware := func(ginContext *gin.Context) {
		fmt.Println("defaultMiddleware Before Next:", ginContext.FullPath())
		ginContext.Header("X-Request-Id", uuid.New().String())
		ginContext.Next()
		fmt.Println("defaultMiddleware After Next:", ginContext.FullPath())
		ginContext.Header("X-Response-Id", uuid.New().String())
	}
	ginEngine.Use(defaultMiddleware)
	ginEngine.GET("/middleware", func(ginContext *gin.Context) {
		fmt.Println("1 Middle Before Next")
		ginContext.Next()
		fmt.Println("1 Middle After Next")
	}, func(ginContext *gin.Context) {
		fmt.Println("2 Middle Before Next")
		// 在第二個 middleware 執行 abort
		// ginContext.AbortWithStatus(http.StatusNotModified)
		ginContext.Next()
		fmt.Println("2 Middle After Next")
	}, func(ginContext *gin.Context) {
		fmt.Println("3 Middle Before Next")
		ginContext.Next()
		fmt.Println("3 Middle After Next")
	}, func(ginContext *gin.Context) {
		ginContext.String(http.StatusOK, "middleware")
	})
}
