package main

import (
	"crypto/tls"
	"demo.golang/controller"
	"demo.golang/sample"
	"demo.golang/singleton"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"time"
)

func backend() {
	fmt.Println(singleton.SingletonConfiguration.Mode)

	// gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
	// ginLogFile, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(ginLogFile)
	// gin.DefaultWriter = io.MultiWriter(AppLogrus.Writer())

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		fmt.Printf("endpoint [%v] [%v] [%v] [%v]\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	ginEngine := gin.Default()
	pprof.Register(ginEngine)
	// ginEngine.SetTrustedProxies([]string{"192.168.1.2"})
	// cors
	ginEngine.Use(cors.Default())
	// ginEngine.Use(cors.New(CorsConfig()))
	// Global middleware
	ginEngine.Use(func(ginContext *gin.Context) {
		ginContext.Request.Header.Add("scopeId", uuid.NewString())
		ginContext.Next()
	})
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	// ginEngine.Use(gin.Logger())
	ginEngine.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] [%s] [%36s] [%s] [%s] [%3d] [%s] [%s]\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Request.Header.Get("uuid"),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.ErrorMessage,
		)
	}))
	// ginEngine.Use(gin.Recovery())
	ginEngine.Use(gin.CustomRecovery(func(ginContext *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			ginContext.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		ginContext.AbortWithStatus(http.StatusInternalServerError)
	}))
	sample.App(ginEngine)
	controller.Router(ginEngine)
	// static
	ginEngine.Static("/assets", "./assets")
	// ginEngine.Static("/assets", "./assets")
	// ginEngine.StaticFS("/more_static", http.Dir("my_file_system"))
	ginEngine.StaticFile("/favicon.ico", "./favicon.ico")
	// ginEngine.StaticFileFS("/more_favicon.ico", "more_favicon.ico", http.Dir("my_file_system"))

	if _, err := tls.LoadX509KeyPair("", ""); err != nil {
		fmt.Println(err)
	}
	// if _, err := tls.LoadX509KeyPair(common.AppConfiguration.Mkcert.Cert, common.AppConfiguration.Mkcert.Key); err != nil {
	// 	fmt.Println(err)
	// }
	// gin.Engine
	ginEngine.Run(":80")
	// ginEngine.RunTLS(":443", common.AppConfiguration.Mkcert.Cert, common.AppConfiguration.Mkcert.Key)
	// http.ListenAndServe
	// if err := http.ListenAndServe(":80", ginEngine); err != nil {
	// 	fmt.Println(err)
	// }
	// if err := http.ListenAndServeTLS(":443", AppConfiguration.Mkcert.Cert, AppConfiguration.Mkcert.Key, ginEngine); err != nil {
	// 	fmt.Println(err)
	// }
	// Custom HTTP configuration
	// server := &http.Server{
	// 	Handler:        ginEngine,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// server.Addr = ":80"
	// if err := server.ListenAndServe(); err != nil {
	// 	fmt.Println(err)
	// }
	// server.Addr = ":443"
	// if err := server.ListenAndServeTLS(AppConfiguration.Mkcert.Cert, AppConfiguration.Mkcert.Key); err != nil {
	// 	fmt.Println(err)
	// }
}

// func CorsConfig(conf *cors.Config) cors.Config {
//     // corsConf := cors.DefaultConfig()
//     corsConf := cors.Config{
//         MaxAge:                 12 * time.Hour,
//         AllowBrowserExtensions: true,
//     }
//     if mode.IsDev() {
//         // 在開發環境時，允許所有 origins、所有 methods 和多數的 headers
//         corsConf.AllowAllOrigins = true
//         corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
//         corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
//             "Connection", "Accept-Encoding", "Accept-Language", "Host"}
//     } else {
//         // 在正式環境時則根據設定檔調整
//         compiledOrigins := compileAllowedCORSOrigins(conf.Server.Cors.AllowOrigins)
//         corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
//         corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Origin",
//             "Connection", "Accept-Encoding", "Accept-Language", "Host"}
//         corsConf.AllowOrigins = []string{"https://www.example.com"}
//     }
//     return corsConf
// }
