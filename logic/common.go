package logic

import (
	"context"
	"fmt"
)

type ICommonLogic interface {
	CommonRun(context.Context) error
}

func (i *Logic) CommonRun(ctx context.Context) error {
	fmt.Println("CommonRun")
	return nil
}
