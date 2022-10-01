package logic

import (
	"context"
	"demo.golang/repository"
)

type ILogic interface {
	Run(context.Context) error
	ICommonLogic
	IDefaultLogic
}
type Logic struct {
	Repository repository.IRepository
}

func NewLogic() ILogic {
	return &Logic{
		Repository: repository.NewRepository(),
	}
}

func (i *Logic) Run(ctx context.Context) error {
	return i.Repository.WithTransaction(ctx, func(ctx context.Context, repo repository.IRepository) error {
		return repo.Run(ctx)
	})
	//return i.Repository.Run(ctx)
}
