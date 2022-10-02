package repository

import (
	"context"
	"database/sql"
	"demo.golang/dto"
	"demo.golang/entity"
	"fmt"
)

type IDefaultRepository interface {
	DefaultRun(context.Context) error
	DefaultQueryDto(context.Context) (*[]dto.DefaultPersonDto, error)
	DefaultQueryAdvanced(context.Context) (*[]entity.Person, error)
	DefaultQueryRow(context.Context) (*sql.Rows, error)
}

func (i *Repository) DefaultRun(ctx context.Context) error {
	fmt.Println("DefaultRun")
	return nil
}
func (i *Repository) DefaultQueryDto(ctx context.Context) (*[]dto.DefaultPersonDto, error) {
	data := &[]dto.DefaultPersonDto{}
	err := i.Db.WithContext(ctx).
		Raw("SELECT row as a, id as b, name as c FROM person").
		Scan(data).
		Error
	return data, err
}
func (i *Repository) DefaultQueryAdvanced(ctx context.Context) (*[]entity.Person, error) {
	data := &[]entity.Person{}
	err := i.Db.WithContext(ctx).Model(entity.Person{}).
		// Select().
		Where("row > ?", 0).
		// Where("row > ?", 0).
		// Or("row > ?", 0).
		Order("row").
		// Order("row ASC").
		// Order("row DESC").
		Limit(2).
		Find(data).
		Error
	return data, err
}
func (i *Repository) DefaultQueryRow(ctx context.Context) (*sql.Rows, error) {
	return i.Db.WithContext(ctx).
		Raw("SELECT row as a, id as b, name as c FROM person").
		Rows()
}
