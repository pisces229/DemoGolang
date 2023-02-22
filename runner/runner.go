package runner

import (
	"context"
	"demo.golang/logic"
)

type IRunner interface {
	DefaultRunner(context.Context) error
}

type Runner struct {
	Logic logic.ILogic
}

func NewRunner() IRunner {
	return &Runner{
		Logic: logic.NewLogic(),
	}
}
