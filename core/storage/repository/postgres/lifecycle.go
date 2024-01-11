package postgres

import (
	"context"
	"gosi/core/domain/lifecycle"
	"sync"

	"golang.org/x/sync/errgroup"
)

type disctionaryData struct {
	lifecycleStates map[int]lifecycle.LifecycleState
	lifecycles      map[int]lifecycle.Lifecycle
}

type lifecycleRepository struct {
	lockDb     *sync.RWMutex
	db         IPostgresDatabase
	dictionary disctionaryData

	eg  *errgroup.Group
	ctx context.Context
}

func (repo *lifecycleRepository) loadLifecycles() {
	//TODO: TO IMPLEMENT
	// var (
	// 	lifecycleStatesRows []lifecycle.LifecycleStateRow
	// 	lifecyclesRows      []lifecycle.LifecycleRow
	// )
	// err := repo.db.NewSelect().Model(&lifecycleStatesRows).Scan(repo.ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, row := range lifecycleStatesRows {
	// 	repo.dictionary.lifecycleStates[row.Id] = NewLifecycleState(row.Id, row.Name)
	// }
	//
	// err = repo.db.NewSelect().Model(&lifecyclesRows).Scan(repo.ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, row := range lifecyclesRows {
	// 	repo.dictionary.lifecycles[row.Id] = NewLifeCycleBuilder(row.Id, row.Name, repo.dictionary.lifecycleStates[row.StartStateId]).Build()
	// }
}
