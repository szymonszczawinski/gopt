package auth

import (
	"context"
	"gopt/core/domain/auth"
	"gopt/core/storage/repository/postgres"
	"sync"

	"golang.org/x/sync/errgroup"
)

type authRepository struct {
	lockDb *sync.RWMutex
	db     postgres.IPostgresDatabase
	eg     *errgroup.Group
	ctx    context.Context
}

func NewAuthRepository(eg *errgroup.Group, ctx context.Context, db postgres.IPostgresDatabase) auth.IAuthRepository {
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
