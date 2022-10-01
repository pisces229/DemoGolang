package sample

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func demoRequestUpload(ginEngine *gin.Engine) {
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	ginEngine.MaxMultipartMemory = 8 << 20 // 8 MiB
	dst := "./temp"
	ginEngine.POST("/upload", func(ginContext *gin.Context) {
		file, err := ginContext.FormFile("file")
		if err == nil {
			fmt.Println(file)
			ginContext.SaveUploadedFile(file, filepath.Join(dst, uuid.New().String()))
			ginContext.Status(http.StatusOK)
		} else {
			ginContext.Status(http.StatusBadRequest)
		}
	})
	ginEngine.POST("/uploads", func(ginContext *gin.Context) {
		form, err := ginContext.MultipartForm()
		if err == nil {
			files := form.File["file"]
			for _, file := range files {
				fmt.Println(file)
				ginContext.SaveUploadedFile(file, filepath.Join(dst, uuid.New().String()))
			}
			ginContext.Status(http.StatusOK)
		} else {
			ginContext.Status(http.StatusBadRequest)
		}
	})
	ginEngine.POST("/form", func(ginContext *gin.Context) {
		form, err := ginContext.MultipartForm()
		if err == nil {
			for key := range form.Value {
				values := form.Value[key]
				fmt.Println(key, values)
			}
			for key := range form.File {
				files := form.File[key]
				for _, file := range files {
					fmt.Println(key, file)
				}
			}
			ginContext.Status(http.StatusOK)
		} else {
			ginContext.Status(http.StatusBadRequest)
		}
	})
}
