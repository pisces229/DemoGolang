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

func (i *Logic) Run(ctx context.Context) (err error) {
	// Auto
	return i.Repository.Run(ctx)

	// WithTransaction
	// return i.Repository.WithTransaction(ctx, func(ctx context.Context, repo repository.IRepository) error {
	// 	return repo.Run(ctx)
	// })

	// Transaction
	// txRepository := i.Repository.BeginTransaction(ctx)
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		err = fmt.Errorf("transaction painc: %s, rollback error: %s", r, txRepository.Rollback(ctx).Error())
	// 	}
	// }()
	// err = txRepository.Run(ctx)
	// if err == nil {
	// 	if err := txRepository.Commit(ctx); err != nil {
	// 		return err
	// 	}
	// } else {
	// 	if err := txRepository.Rollback(ctx); err != nil {
	// 		return err
	// 	}
	// }
	// return
}
