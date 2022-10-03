package repository

import (
	"context"
	"fmt"
	"time"

	"demo.golang/entity"
	"demo.golang/singleton"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IRepository interface {
	WithTransaction(context.Context, func(ctx context.Context, repo IRepository) error) error
	BeginTransaction(context.Context) IRepository
	Commit(context.Context) error
	Rollback(context.Context) error
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
		Db: singleton.SingletonDatabase,
	}
}

// WithTransaction ...
func (i *Repository) WithTransaction(ctx context.Context, fn func(ctx context.Context, repo IRepository) error) (err error) {
	err = i.Db.WithContext(ctx).Transaction(func(transaction *gorm.DB) error {
		txRepo := &Repository{
			Db: transaction,
		}
		return fn(ctx, txRepo)
	})
	return err
}
func (i *Repository) BeginTransaction(ctx context.Context) IRepository {
	return &Repository{
		Db: i.Db.WithContext(ctx).Begin(),
	}
}
func (i *Repository) Commit(ctx context.Context) error {
	return i.Db.WithContext(ctx).Commit().Error
}
func (i *Repository) Rollback(ctx context.Context) error {
	return i.Db.WithContext(ctx).Rollback().Error
}

// Run ...
func (i *Repository) Run(ctx context.Context) error {
	personId := uuid.New().String()
	if err := i.DatabaseCreate(ctx, &entity.Person{
		Id:       personId,
		Name:     uuid.New().String(),
		Age:      1,
		Birthday: time.Now(),
		Remark:   uuid.New().String(),
	}); err != nil {
		fmt.Println(err)
		return err
	}

	{
		persons := []entity.Person{
			{
				Id:       uuid.New().String(),
				Name:     uuid.New().String(),
				Age:      1,
				Birthday: time.Now(),
				Remark:   uuid.New().String(),
			},
			{
				Id:       uuid.New().String(),
				Name:     uuid.New().String(),
				Age:      1,
				Birthday: time.Now(),
				Remark:   uuid.New().String(),
			},
		}
		if err := i.DatabaseCreate(ctx, persons); err != nil {
			fmt.Println(err)
			return err
		}
	}

	{
		persons := &[]entity.Person{}
		if err := i.DatabaseFind(ctx, entity.Person{}, persons, []int{}); err != nil {
			fmt.Println(err)
			return err
		} else {
			for _, person := range *persons {
				if err := i.DatabaseModify(ctx, &person, entity.Person{Remark: uuid.New().String()}); err != nil {
					fmt.Println(err)
					return err
				}
			}
			if err := i.DatabaseRemove(ctx, persons); err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	if err := i.DatabaseCreate(ctx, &entity.Person{
		Id:       personId,
		Name:     uuid.New().String(),
		Age:      1,
		Birthday: time.Now(),
		Remark:   uuid.New().String(),
	}); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
