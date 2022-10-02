package runner

import (
	"context"
	"demo.golang/logic"
)

type IRunner interface {
	DefaultRunner() error
}

type Runner struct {
	Context context.Context
	Logic   logic.ILogic
}

func NewRunner() IRunner {
	return &Runner{
		Context: context.TODO(),
		Logic:   logic.NewLogic(),
	}
}
