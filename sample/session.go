package sample

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func demoSession(ginEngine *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	// ginEngine.Use(sessions.Sessions("session-name", store))
	ginEngine.Use(sessions.SessionsMany([]string{"count", "first", "second"}, store))
	ginEngine.GET("/session", func(ginContext *gin.Context) {
		// session := sessions.Default(c)
		session := sessions.DefaultMany(ginContext, "count")
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		ginContext.JSON(200, gin.H{"count": count})
	})
}
