package auth

import (
	"context"
	"gosi/core/storage/sql"
	"gosi/coreapi/service"
	"sync"

	"golang.org/x/sync/errgroup"
)

type IAuthRepository interface {
	service.IComponent
}

type authRepository struct {
	lockDb *sync.RWMutex
	db     sql.IPostgresDatabase
	eg     *errgroup.Group
	ctx    context.Context
}

func NewAuthRepository(eg *errgroup.Group, ctx context.Context, db sql.IPostgresDatabase) IAuthRepository {
	instance := authRepository{
		lockDb: &sync.RWMutex{},
		db:     db,
		eg:     eg,
		ctx:    ctx,
	}
	return instance
}

func (repo authRepository) StartComponent() {
}
