package sample

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func demoResponse(ginEngine *gin.Engine) {
	ginEngine.GET("/string", func(ginContext *gin.Context) {
		ginContext.String(http.StatusOK, "golang")
	})
	ginEngine.GET("/json", func(ginContext *gin.Context) {
		type DataStruct struct {
			First  string `json:"first"`
			Second string `json:"second"`
		}
		jsonBytes := []byte(`[{"first" : "aaa1嗨" , "second" : "bbb1"},{"first" : "aaa2" , "second" : "bbb2"}]`)
		var dataStruct []DataStruct
		json.Unmarshal(jsonBytes, &dataStruct)
		fmt.Println(dataStruct)
		jsonData, _ := json.Marshal(dataStruct)
		fmt.Println(string(jsonData))
		// ginContext.AsciiJSON(http.StatusOK, gin.H{
		// 	"language": "Golang",
		// 	"message":  "Go語言",
		// })
		ginContext.JSON(http.StatusOK, gin.H{
			"language": "Golang",
			"message":  "Go語言",
		})
	})
	ginEngine.GET("/json1", func(ginContext *gin.Context) {
		type DataStruct struct {
			First  string `json:"first"`
			Second string `json:"second"`
		}
		result := &DataStruct{First: "Golang", Second: "Go語言"}
		// ginContext.AsciiJSON(http.StatusOK, result)
		ginContext.JSON(http.StatusOK, result)
	})
	ginEngine.GET("/json2", func(ginContext *gin.Context) {
		type DataStruct struct {
			First  string `json:"first"`
			Second string `json:"second"`
		}
		var result []DataStruct
		result = append(result, DataStruct{First: "Golang1", Second: "Go語言1"})
		result = append(result, DataStruct{First: "Golang2", Second: "Go語言2"})
		// ginContext.AsciiJSON(http.StatusOK, result)
		ginContext.JSON(http.StatusOK, result)
	})
	ginEngine.GET("/json3", func(ginContext *gin.Context) {
		type DataStruct struct {
			First  string `json:"first"`
			Second string `json:"second"`
		}
		type RootStruct struct {
			Success bool         `json:"success"`
			Message string       `json:"message"`
			Data    []DataStruct `json:"data"`
		}
		var data []DataStruct
		data = append(data, DataStruct{First: "Golang1", Second: "Go語言1"})
		data = append(data, DataStruct{First: "Golang2", Second: "Go語言2"})
		result := &RootStruct{Success: true, Message: "Success", Data: data}
		// ginContext.AsciiJSON(http.StatusOK, result)
		ginContext.JSON(http.StatusOK, result)
	})
	ginEngine.GET("/download", func(ginContext *gin.Context) {
		// ginContext.File("temp/gin.png")
		// ginContext.File("temp/gin.txt")
		ginContext.File("temp/gin.rtf")
	})
	ginEngine.GET("/download1", func(ginContext *gin.Context) {
		// file := http.FileServer(http.Dir("/temp"))
		// if err == nil {
		// 	ginContext.FileFromFS("temp/gin.rtf", file)
		// } else {
		// 	ginContext.Status(http.StatusBadRequest)
		// }
	})
	ginEngine.GET("/download2", func(ginContext *gin.Context) {
		// ginContext.FileAttachment("temp/gin.png", "donwload.png")
		// ginContext.FileAttachment("temp/gin.txt", "donwload.txt")
		ginContext.FileAttachment("temp/gin.rtf", "donwload.rtf")
	})
	ginEngine.GET("/download3", func(ginContext *gin.Context) {
		data, err := os.ReadFile("temp/gin.png")
		fmt.Println("Error:", err)
		if err == nil {
			ginContext.Data(http.StatusOK, "application/octet-stream", data)
		} else {
			ginContext.Status(http.StatusBadRequest)
		}
	})
	ginEngine.GET("/download4", func(ginContext *gin.Context) {
		data, err := os.ReadFile("temp/gin.png")
		fmt.Println("Error:", err)
		if err == nil {
			reader := bytes.NewReader(data)
			contentLength := int64(len(data))
			contentType := "application/octet-stream"
			extraHeaders := map[string]string{
				"Content-Disposition": `attachment; filename="donwload.png"`,
			}
			ginContext.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
		} else {
			ginContext.Status(http.StatusBadRequest)
		}
	})
}
