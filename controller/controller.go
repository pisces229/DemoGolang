package controller

import (
	"demo.golang/logic"
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
