package domain

import (
	"context"

	"gorm.io/gorm"
)

type CrudRepository interface {
	GetAll(ctx context.Context) ([]Items, error)
	GetById(ctx context.Context, request RequestGet) (Items, error)
	PostItem(ctx context.Context, request Items) error
	PutItem(ctx context.Context, request Items) error
	DeleteItem(ctx context.Context, request RequestGet) error
}
type CrudUsecase interface {
	GetAll(ctx context.Context) ([]Items, error)
	GetById(ctx context.Context, request RequestGet) (Items, error)
	PostItem(ctx context.Context, request Items) error
	PutItem(ctx context.Context, request Items) error
	DeleteItem(ctx context.Context, request RequestGet) error
}

type RequestGet struct {
	Id uint `json:"id"`
}

type Items struct {
	gorm.Model
	Detail string `json:"detail" gorm:"column:detail;type:varchar"`
}
