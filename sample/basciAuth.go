package sample

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func demoBasicAuth(ginEngine *gin.Engine) {
	data := gin.H{
		"guest": gin.H{"email": "guest@mail.com", "phone": "12345"},
		"admin": gin.H{"email": "admin@mail.com", "phone": "67890"},
	}
	authorized := ginEngine.Group("/admin",
		gin.BasicAuth(gin.Accounts{
			"guest": "0000",
			"admin": "1234",
		}))
	// endpoint: /admin/secrets
	authorized.GET("/login", func(ginContext *gin.Context) {
		user := ginContext.MustGet(gin.AuthUserKey).(string)
		if data, ok := data[user]; ok {
			ginContext.JSON(http.StatusOK, gin.H{"user": user, "data": data})
		} else {
			ginContext.JSON(http.StatusOK, gin.H{"user": user, "data": "NO SECRET :("})
		}
	})
}
