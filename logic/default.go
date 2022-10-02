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
	DefaultQueries(context.Context) error
}

func (i *Logic) DefaultRun(ctx context.Context) error {
	fmt.Println("DefaultRun")
	return i.Repository.DefaultRun(ctx)
}
func (i *Logic) DefaultQuery(ctx context.Context) error {
	persons := &[]entity.Person{}
	err := i.Repository.DatabaseFind(ctx, entity.Person{}, persons, []int{})
	fmt.Println(persons)
	return err
}
func (i *Logic) DefaultCreate(ctx context.Context) error {
	person := &entity.Person{
		Id:       uuid.New().String(),
		Name:     uuid.New().String(),
		Age:      1,
		Birthday: time.Now(),
		Remark:   uuid.New().String(),
	}
	err := i.Repository.DatabaseCreate(ctx, person)
	fmt.Println(person)
	return err
}
func (i *Logic) DefaultModify(ctx context.Context) error {
	persons := &[]entity.Person{}
	err := i.Repository.DatabaseFind(ctx, entity.Person{}, persons, []int{})
	for _, person := range *persons {
		i.Repository.DatabaseModify(ctx, &person, entity.Person{Remark: uuid.New().String()})
	}
	return err
}
func (i *Logic) DefaultRemove(ctx context.Context) error {
	persons := &[]entity.Person{}
	err := i.Repository.DatabaseFind(ctx, entity.Person{}, persons, []int{})
	i.Repository.DatabaseRemove(ctx, persons)
	return err
}
func (i *Logic) DefaultTransaction(ctx context.Context) error {
	return i.Repository.WithTransaction(ctx,
		func(ctx context.Context, repository repository.IRepository) error {
			person := &entity.Person{
				Id:       uuid.New().String(),
				Name:     uuid.New().String(),
				Age:      1,
				Birthday: time.Now(),
				Remark:   uuid.New().String(),
			}
			if err := i.Repository.DatabaseCreate(ctx, person); err != nil {
				return err
			}
			person.Row = 0
			if err := i.Repository.DatabaseCreate(ctx, person); err != nil {
				return err
			}
			return nil
		})
}
func (i *Logic) DefaultQueries(ctx context.Context) error {
	if data, err := i.Repository.DefaultQueryDto(ctx); err != nil {
		return err
	} else {
		fmt.Println("DefaultQueryDto", data)
	}
	if data, err := i.Repository.DefaultQueryAdvanced(ctx); err != nil {
		return err
	} else {
		fmt.Println("DefaultQueryAdvanced", data)
	}
	if rows, err := i.Repository.DefaultQueryRow(ctx); err != nil {
		return err
	} else {
		defer func() {
			if err := rows.Close(); err != nil {
				fmt.Println(err)
			}
		}()
		var a int32
		var b string
		var c string
		for rows.Next() {
			if err := rows.Scan(&a, &b, &c); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("DefaultQueryRow", a, b, c)
			}
		}
	}
	return nil
}
