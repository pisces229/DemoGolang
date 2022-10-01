package controller

import (
	"demo.golang/logic"
	"github.com/gin-gonic/gin"
)

type IController interface {
	IDefaultController
}

type Controller struct {
	Logic logic.ILogic
}

func NewController() IController {
	return &Controller{
		Logic: logic.NewLogic(),
	}
}

func Router(ginEngine *gin.Engine) {
	controller := NewController()
	controller.DefaultRouter(ginEngine)
}
