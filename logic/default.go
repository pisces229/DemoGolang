package logic

import (
	"context"
	"demo.golang/entity"
	"demo.golang/repository"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type IDefaultLogic interface {
	DefaultRun(context.Context) error
	DefaultQuery(context.Context) error
	DefaultCreate(context.Context) error
	DefaultModify(context.Context) error
	DefaultRemove(context.Context) error
	DefaultTransaction(context.Context) error
}

func (i *Logic) DefaultRun(ctx context.Context) error {
	fmt.Println("DefaultRun")
	return i.Repository.DefaultRun(ctx)
}
func (i *Logic) DefaultQuery(ctx context.Context) error {
	customers := &[]entity.Customer{}
	err := i.Repository.DatabaseFind(ctx, entity.Customer{}, customers, []int{})
	fmt.Println(customers)
	return err
}
func (i *Logic) DefaultCreate(ctx context.Context) error {
	customer := &entity.Customer{
		Id:       "1",
		Name:     uuid.New().String(),
		Age:      1,
		Birthday: time.Now(),
		Remark:   uuid.New().String(),
	}
	err := i.Repository.DatabaseCreate(ctx, customer)
	fmt.Println(customer)
	return err
}
func (i *Logic) DefaultModify(ctx context.Context) error {
	customers := &[]entity.Customer{}
	err := i.Repository.DatabaseFind(ctx, entity.Customer{}, customers, []int{1})
	for _, customer := range *customers {
		i.Repository.DatabaseModify(ctx, customer, entity.Customer{Remark: uuid.New().String()})
	}
	return err
}
func (i *Logic) DefaultRemove(ctx context.Context) error {
	customers := &[]entity.Customer{}
	err := i.Repository.DatabaseFind(ctx, entity.Customer{}, customers, []int{1})
	i.Repository.DatabaseModify(ctx, customers, entity.Customer{Remark: uuid.New().String()})
	return err
}
func (i *Logic) DefaultTransaction(ctx context.Context) error {
	return i.Repository.WithTransaction(ctx,
		func(ctx context.Context, repository repository.IRepository) error {
			customer := &entity.Customer{
				Id:       "1",
				Name:     uuid.New().String(),
				Age:      1,
				Birthday: time.Now(),
				Remark:   uuid.New().String(),
			}
			if err := i.Repository.DatabaseCreate(ctx, customer); err != nil {
				return err
			}
			if err := i.Repository.DatabaseCreate(ctx, customer); err != nil {
				return err
			}
			return nil
		})
}
