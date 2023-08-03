package auth

import (
	"context"
	"gosi/coreapi/service"
	"gosi/coreapi/storage"
	"sync"

	"golang.org/x/sync/errgroup"
)

type IAuthRepository interface {
	service.IComponent
}

type authRepository struct {
	lockDb *sync.RWMutex
	db     storage.IBunDatabase

	eg  *errgroup.Group
	ctx context.Context
}

func NewAuthRepository(eg *errgroup.Group, ctx context.Context, db storage.IBunDatabase) IAuthRepository {
	instance := authRepository{
		lockDb: &sync.RWMutex{},
		db:     db,
		eg:     eg,
		ctx:    ctx,
	}
	return instance
}
func (self authRepository) StartComponent() {
}
