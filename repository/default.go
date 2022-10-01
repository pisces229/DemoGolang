package repository

import (
	"context"
	"fmt"
)

type IDefaultRepository interface {
	DefaultRun(context.Context) error
	DefaultQuery(context.Context) error
}

func (i *Repository) DefaultRun(ctx context.Context) error {
	fmt.Println("DefaultRun")
	return nil
}
func (i *Repository) DefaultQuery(ctx context.Context) error {
	return nil
}
