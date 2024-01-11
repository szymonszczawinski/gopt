package postgres

import (
	"context"
	"gosi/core/domain/auth"
	"sync"

	"golang.org/x/sync/errgroup"
)

type authRepository struct {
	lockDb *sync.RWMutex
	db     IPostgresDatabase
	eg     *errgroup.Group
	ctx    context.Context
}

func NewAuthRepository(eg *errgroup.Group, ctx context.Context, db IPostgresDatabase) auth.IAuthRepository {
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
