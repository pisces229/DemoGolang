package repository

import "context"

type ICommonRepository interface {
	CommonRun(context.Context) error
}

func (i *Repository) CommonRun(ctx context.Context) error {
	return nil
}
