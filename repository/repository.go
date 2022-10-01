package repository

import (
	"context"
	"demo.golang/entity"
	"demo.golang/singleton"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type IRepository interface {
	WithTransaction(context.Context, func(ctx context.Context, repo IRepository) error) error
	IDatabase
	ICommonRepository
	IDefaultRepository
	Run(context.Context) error
}
type Repository struct {
	Db *gorm.DB
}

func NewRepository() IRepository {
	return &Repository{
		Db: singleton.AppDatabase,
	}
}

// WithTransaction ...
func (i *Repository) WithTransaction(ctx context.Context, fn func(ctx context.Context, repo IRepository) error) error {
	err := i.Db.WithContext(ctx).Transaction(func(transaction *gorm.DB) error {
		txRepo := &Repository{
			Db: transaction,
		}
		return fn(ctx, txRepo)
	})
	return err
}

// Run ...
func (i *Repository) Run(ctx context.Context) error {
	if err := i.DatabaseCreate(ctx, &entity.Customer{
		Id:       "1",
		Name:     uuid.New().String(),
		Age:      1,
		Birthday: time.Now(),
		Remark:   uuid.New().String(),
	}); err != nil {
		fmt.Println(err)
		return err
	}

	{
		customers := []entity.Customer{
			entity.Customer{
				Id:       uuid.New().String(),
				Name:     uuid.New().String(),
				Age:      1,
				Birthday: time.Now(),
				Remark:   uuid.New().String(),
			},
			entity.Customer{
				Id:       uuid.New().String(),
				Name:     uuid.New().String(),
				Age:      1,
				Birthday: time.Now(),
				Remark:   uuid.New().String(),
			},
		}
		if err := i.DatabaseCreate(ctx, customers); err != nil {
			fmt.Println(err)
			return err
		}
	}

	{
		customers := &[]entity.Customer{}
		if err := i.DatabaseFind(ctx, entity.Customer{}, customers, []int{}); err != nil {
			fmt.Println(err)
			return err
		} else {

			for _, customer := range *customers {
				i.DatabaseModify(ctx, customer, entity.Customer{Remark: uuid.New().String()})
			}

			i.DatabaseRemove(ctx, customers)
		}
	}

	//if err := i.DatabaseCreate(ctx, &entity.Customer{
	//	Id:       "1",
	//	Name:     uuid.New().String(),
	//	Age:      1,
	//	Birthday: time.Now(),
	//	Remark:   uuid.New().String(),
	//}); err != nil {
	//	fmt.Println(err)
	//	return err
	//}

	return nil
}
