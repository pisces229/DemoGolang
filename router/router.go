package router

import (
	"demo.golang/controller"
	"demo.golang/router/sample"
	"github.com/gin-gonic/gin"
)

type IRouter interface {
	Router()
	IDefaultRouter
}

type Router struct {
	Engine     *gin.Engine
	Controller controller.IController
}

func (i *Router) Router() {
	sample.Middleware(i.Engine)
	//sample.Session(i.Engine)
	//sample.Cookie(i.Engine)
	//sample.Jwt(i.Engine)
	//sample.BasicAuth(i.Engine)
	//sample.RequestGet(i.Engine)
	//sample.RequestPost(i.Engine)
	//sample.RequestUpload(i.Engine)
	//sample.Response(i.Engine)
	//sample.HtmlRendering(i.Engine)
	//sample.Redirect(i.Engine)
	i.DefaultRouter()
}

func NewRouter(engine *gin.Engine) IRouter {
	return &Router{
		Engine:     engine,
		Controller: controller.NewController(),
	}
}
