package sample

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequestGet(ginEngine *gin.Engine) {
	ginEngine.POST("/result", func(ginContext *gin.Context) {
		ginContext.Status(http.StatusOK)
	})
	// Parameters In Path
	// 使用 :<filed> 可以定義動態路由（只能匹配到 / 以前）
	// 使用 *<filed> 可以定義動態路由（可以匹配到 / 以後）
	ginEngine.GET("/ParametersInPath", func(ginContext *gin.Context) {
		fmt.Printf("[%s]\n", ginContext.FullPath())
		ginContext.Status(http.StatusOK)
	})
	// 不會匹配到 /ParametersInPath/ 或 /ParametersInPath
	ginEngine.GET("/ParametersInPath/:first", func(ginContext *gin.Context) {
		fmt.Printf("[%s][%s]\n", ginContext.FullPath(), ginContext.Param("first"))
		ginContext.Status(http.StatusOK)
	})
	// 然而，這將會匹配到 /ParametersInPath/john/ 和 /ParametersInPath/john/send
	ginEngine.GET("/ParametersInPath/:first/*second", func(ginContext *gin.Context) {
		fmt.Printf("[%s][%s][%s]\n", ginContext.FullPath(), ginContext.Param("first"), ginContext.Param("second"))
		ginContext.Status(http.StatusOK)
	})
	// Querystring In Parameters
	ginEngine.GET("/QuerystringInParameters", func(ginContext *gin.Context) {
		// ginContext.Query("name")
		// ginContext.Request.URL.Query().Get("name")
		fmt.Printf("[%s][%s][%s]\n", ginContext.FullPath(), ginContext.DefaultQuery("first", "Unknow"), ginContext.Query("second"))
		ginContext.Status(http.StatusOK)
	})
	ginEngine.GET("/BindUri/:first/:second", func(ginContext *gin.Context) {
		type DataStruct struct {
			First  string `uri:"first"`
			Second string `uri:"second"`
		}
		var data DataStruct
		// context.BindUri(&data)
		// fmt.Println("Bind:", data)
		err := ginContext.ShouldBindUri(&data)
		fmt.Println("ShouldBindUri:", data)
		fmt.Println("ShouldBindUri Error:", err)
		if err == nil {
			ginContext.Status(http.StatusOK)
		} else {
			ginContext.Status(http.StatusBadRequest)
		}
	})
	ginEngine.GET("/BindUri/Validation/:first/:second", func(ginContext *gin.Context) {
		type DataStruct struct {
			First  string `uri:"first" binding:"required"`
			Second string `uri:"second" binding:"required"`
		}
		data := &DataStruct{}
		err := ginContext.ShouldBindUri(data)
		fmt.Println("ShouldBindUri:", *data)
		fmt.Println("ShouldBindUri Error:", err)
		if err == nil {
			ginContext.Status(http.StatusOK)
		} else {
			ginContext.Status(http.StatusBadRequest)
		}
	})
	ginEngine.GET("/Bind", func(ginContext *gin.Context) {
		type DataStruct struct {
			First  string `form:"first"`
			Second string `form:"second"`
		}
		data := &DataStruct{}
		// ginContext.Bind(&data)
		// fmt.Println("Bind:", data)
		err := ginContext.ShouldBind(data)
		fmt.Println("ShouldBind:", *data)
		fmt.Println("ShouldBind Error:", err)
		if err == nil {
			ginContext.Status(http.StatusOK)
		} else {
			ginContext.Status(http.StatusBadRequest)
		}
	})
	ginEngine.GET("/Bind/Validation", func(ginContext *gin.Context) {
		type DataStruct struct {
			First  string `form:"first" binding:"required"`
			Second string `form:"second" binding:"required"`
		}
		data := &DataStruct{}
		err := ginContext.ShouldBind(data)
		fmt.Println("ShouldBind:", *data)
		fmt.Println("ShouldBind Error:", err)
		if err == nil {
			ginContext.Status(http.StatusOK)
		} else {
			ginContext.Status(http.StatusBadRequest)
		}
	})
	ginEngine.GET("/BindQuery", func(ginContext *gin.Context) {
		type DataStruct struct {
			First  string `form:"first"`
			Second string `form:"second"`
		}
		data := &DataStruct{}
		// ginContext.BindQuery(&data)
		// fmt.Println("Bind:", data)
		err := ginContext.ShouldBindQuery(data)
		fmt.Println("BindQuery:", *data)
		fmt.Println("BindQuery Error:", err)
		if err == nil {
			ginContext.Status(http.StatusOK)
		} else {
			ginContext.Status(http.StatusBadRequest)
		}
	})
	ginEngine.GET("/BindQuery/Validation", func(ginContext *gin.Context) {
		type DataStruct struct {
			First  string `form:"first" binding:"required"`
			Second string `form:"second" binding:"required"`
		}
		data := &DataStruct{}
		err := ginContext.ShouldBindQuery(data)
		fmt.Println("ShouldBindQuery:", *data)
		fmt.Println("ShouldBindQuery Error:", err)
		if err == nil {
			ginContext.Status(http.StatusOK)
		} else {
			ginContext.Status(http.StatusBadRequest)
		}
	})
}
