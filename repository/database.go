package repository

import (
	"context"
)

type IDatabase interface {
	DatabaseFind(context.Context, interface{}, interface{}, []int) error
	DatabaseCreate(context.Context, interface{}) error
	DatabaseModify(context.Context, interface{}, interface{}) error
	DatabaseRemove(context.Context, interface{}) error
}

func (i *Repository) DatabaseFind(ctx context.Context, model interface{}, data interface{}, rows []int) error {
	return i.Db.WithContext(ctx).Model(model).Find(data, rows).Error
}
func (i *Repository) DatabaseCreate(ctx context.Context, data interface{}) error {
	return i.Db.WithContext(ctx).Create(data).Error
}
func (i *Repository) DatabaseModify(ctx context.Context, data interface{}, values interface{}) error {
	return i.Db.WithContext(ctx).Model(data).Updates(values).Error
}
func (i *Repository) DatabaseRemove(ctx context.Context, data interface{}) error {
	return i.Db.WithContext(ctx).Delete(data).Error
}
