package sample

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TemplateDataStruct struct {
	Title   string
	Message string
}

func HtmlRendering(ginEngine *gin.Engine) {
	// ginEngine.Delims("{[{", "}]}")
	formatAsDate := func(value time.Time) string {
		year, month, day := value.Date()
		return fmt.Sprintf("%d-%02d-%02d", year, month, day)
	}
	ginEngine.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	ginEngine.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	ginEngine.GET("/index", func(ginContext *gin.Context) {
		ginContext.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Index",
			"date":  time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})
	ginEngine.GET("/first", func(ginContext *gin.Context) {
		// data := map[string]interface{}{
		// 	"title":   "First",
		// 	"message": "First Message",
		// }
		// ginContext.HTML(http.StatusOK, "templates/first.tmpl", data)
		data := TemplateDataStruct{Title: "First", Message: "First Message"}
		ginContext.HTML(http.StatusOK, "templates/first.tmpl", data)
	})
	ginEngine.GET("/second", func(ginContext *gin.Context) {
		items := []TemplateDataStruct{}
		items = append(items, TemplateDataStruct{Title: "Second1", Message: "Second1 Message"})
		items = append(items, TemplateDataStruct{Title: "Second2", Message: "Second2 Message"})
		data := map[string]interface{}{
			"Title": "Second",
			"Datas": items,
		}
		ginContext.HTML(http.StatusOK, "templates/second.tmpl", data)
	})
	ginEngine.GET("/third", func(ginContext *gin.Context) {
		article := []map[string]string{}
		article = append(article, map[string]string{
			"Name": "Article 1",
			"Age":  "1",
		})
		article = append(article, map[string]string{
			"Name": "Article 2",
			"Age":  "2",
		})
		title := []map[string]interface{}{}
		title = append(title, map[string]interface{}{
			"Name":    "First",
			"Message": "First Message",
			"Article": article,
		})
		title = append(title, map[string]interface{}{
			"Name":    "Second",
			"Message": "Second Message",
			"Article": article,
		})
		data := map[string]interface{}{
			"Title": title,
		}
		ginContext.HTML(http.StatusOK, "third.tmpl", data)
	})
}
