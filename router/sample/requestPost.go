package sample

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequestPost(ginEngine *gin.Engine) {
	ginEngine.POST("/result", func(ginContext *gin.Context) {
		ginContext.Status(http.StatusOK)
	})
	ginEngine.POST("/json1", func(ginContext *gin.Context) {
		type DataStruct struct {
			First  string `json:"first"`
			Second string `json:"second"`
		}
		data := &DataStruct{}
		// context.Bind(&data)
		// fmt.Println("Bind:", data)
		err := ginContext.ShouldBind(data)
		fmt.Println("Bind:", *data)
		fmt.Println("Bind Error:", err)
		if err == nil {
			ginContext.Status(http.StatusOK)
		} else {
			ginContext.Status(http.StatusBadRequest)
		}
	})
	ginEngine.POST("/json2", func(ginContext *gin.Context) {
		type DataStruct struct {
			First  string `json:"first"`
			Second string `json:"second"`
		}
		type RootStruct struct {
			Title   string       `json:"title"`
			Message string       `json:"message"`
			Data    []DataStruct `json:"data"`
		}
		data := &RootStruct{}
		// ginContext.Bind(&data)
		// fmt.Println("Bind:", data)
		err := ginContext.ShouldBind(data)
		fmt.Println("Bind:", *data)
		fmt.Println("Bind Error:", err)
		if err == nil {
			ginContext.Status(http.StatusOK)
		} else {
			ginContext.Status(http.StatusBadRequest)
		}
	})
	ginEngine.POST("/array1", func(ginContext *gin.Context) {
		type DataStruct struct {
			First  string `json:"first"`
			Second string `json:"second"`
		}
		data := []DataStruct{}
		body, err := ioutil.ReadAll(ginContext.Request.Body)
		fmt.Println(string(body))
		if err == nil {
			err = json.Unmarshal(body, &data)
			fmt.Println(data)
			if err == nil {
				ginContext.Status(http.StatusOK)
			} else {
				ginContext.Status(http.StatusBadRequest)
			}
		} else {
			ginContext.Status(http.StatusBadRequest)
		}
	})
	ginEngine.POST("/array2", func(ginContext *gin.Context) {
		type DataStruct struct {
			First  string `json:"first"`
			Second string `json:"second"`
		}
		data := []DataStruct{}
		// ginContext.Bind(&data)
		// fmt.Println("Bind:", data)
		err := ginContext.ShouldBind(&data)
		fmt.Println("Bind:", data)
		fmt.Println("Bind Error:", err)
		if err == nil {
			ginContext.Status(http.StatusOK)
		} else {
			ginContext.Status(http.StatusBadRequest)
		}
	})
}
