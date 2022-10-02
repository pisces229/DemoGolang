package router

type IDefaultRouter interface {
	DefaultRouter()
}

func (i *Router) DefaultRouter() {
	i.Engine.GET("/default/run", i.Controller.DefaultRun)
	i.Engine.GET("/default/query", i.Controller.DefaultQuery)
	i.Engine.GET("/default/create", i.Controller.DefaultCreate)
	i.Engine.GET("/default/modify", i.Controller.DefaultModify)
	i.Engine.GET("/default/remove", i.Controller.DefaultRemove)
}
